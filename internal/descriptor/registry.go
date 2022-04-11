package descriptor

import (
	"fmt"
	"strings"

	"github.com/golang/glog"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Registry is a registry of information extracted from pluginpb.CodeGeneratorRequest.
type Registry struct {
	// msgs is a mapping from fully-qualified message name to descriptor
	msgs map[string]*Message

	// enums is a mapping from fully-qualified enum name to descriptor
	enums map[string]*Enum

	// files is a mapping from file path to descriptor
	files map[string]*File

	// prefix is a prefix to be inserted to golang package paths generated from proto package names.
	prefix string

	// pkgMap is a user-specified mapping from file path to proto package.
	pkgMap map[string]string

	// pkgAliases is a mapping from package aliases to package paths in go which are already taken.
	pkgAliases map[string]string

	// allowMerge generation one OpenAPI file out of multiple protos
	allowMerge bool

	// mergeFileName target OpenAPI file name after merge
	mergeFileName string

	// includePackageInTags controls whether the package name defined in the `package` directive
	// in the proto file can be prepended to the gRPC service name in the `Tags` field of every operation.
	includePackageInTags bool

	// useJSONNamesForFields if true json tag name is used for generating fields in OpenAPI definitions,
	// otherwise the original proto name is used. It's helpful for synchronizing the OpenAPI definition
	// with gRPC-Gateway response, if it uses json tags for marshaling.
	useJSONNamesForFields bool

	// visibilityRestrictionSelectors is a map of selectors for `google.api.VisibilityRule`s that will be included in the OpenAPI output.
	visibilityRestrictionSelectors map[string]bool

	// useGoTemplate determines whether you want to use GO templates
	// in your protofile comments
	useGoTemplate bool

	// enumsAsInts render enum as integer, as opposed to string
	enumsAsInts bool

	// omitEnumDefaultValue omits default value of enum
	omitEnumDefaultValue bool

	standalone bool
	// proto3OptionalNullable specifies whether Proto3 Optional fields should be marked as x-nullable.
	proto3OptionalNullable bool

	// recursiveDepth sets the maximum depth of a field parameter
	recursiveDepth int
}

// NewRegistry returns a new Registry.
func NewRegistry() *Registry {
	return &Registry{
		msgs:                           make(map[string]*Message),
		enums:                          make(map[string]*Enum),
		files:                          make(map[string]*File),
		pkgMap:                         make(map[string]string),
		pkgAliases:                     make(map[string]string),
		visibilityRestrictionSelectors: make(map[string]bool),
		recursiveDepth:                 1000,
	}
}

func (r *Registry) LoadFromPlugin(gen *protogen.Plugin) error {
	return r.load(gen)
}

func (r *Registry) load(gen *protogen.Plugin) error {
	for filePath, f := range gen.FilesByPath {
		r.loadFile(filePath, f)
	}

	for filePath, f := range gen.FilesByPath {
		if !f.Generate {
			continue
		}
		file := r.files[filePath]
		if err := r.loadServices(file); err != nil {
			return err
		}
	}

	return nil
}

// loadFile loads messages, enumerations and fields from "file".
// It does not loads services and methods in "file".  You need to call
// loadServices after loadFiles is called for all files to load services and methods.
func (r *Registry) loadFile(filePath string, file *protogen.File) {
	pkg := GoPackage{
		Path: string(file.GoImportPath),
		Name: string(file.GoPackageName),
	}
	if r.standalone {
		pkg.Alias = "ext" + strings.Title(pkg.Name)
	}

	if err := r.ReserveGoPackageAlias(pkg.Name, pkg.Path); err != nil {
		for i := 0; ; i++ {
			alias := fmt.Sprintf("%s_%d", pkg.Name, i)
			if err := r.ReserveGoPackageAlias(alias, pkg.Path); err == nil {
				pkg.Alias = alias
				break
			}
		}
	}
	f := &File{
		FileDescriptorProto:     file.Proto,
		GoPkg:                   pkg,
		GeneratedFilenamePrefix: file.GeneratedFilenamePrefix,
	}

	r.files[filePath] = f
	r.registerMsg(f, nil, file.Proto.MessageType)
	r.registerEnum(f, nil, file.Proto.EnumType)
}

func (r *Registry) registerMsg(file *File, outerPath []string, msgs []*descriptorpb.DescriptorProto) {
	for i, md := range msgs {
		m := &Message{
			File:              file,
			Outers:            outerPath,
			DescriptorProto:   md,
			Index:             i,
			ForcePrefixedName: r.standalone,
		}
		for _, fd := range md.GetField() {
			m.Fields = append(m.Fields, &Field{
				Message:              m,
				FieldDescriptorProto: fd,
				ForcePrefixedName:    r.standalone,
			})
		}
		file.Messages = append(file.Messages, m)
		r.msgs[m.FQMN()] = m
		glog.V(1).Infof("register name: %s", m.FQMN())

		var outers []string
		outers = append(outers, outerPath...)
		outers = append(outers, m.GetName())
		r.registerMsg(file, outers, m.GetNestedType())
		r.registerEnum(file, outers, m.GetEnumType())
	}
}

func (r *Registry) registerEnum(file *File, outerPath []string, enums []*descriptorpb.EnumDescriptorProto) {
	for i, ed := range enums {
		e := &Enum{
			File:                file,
			Outers:              outerPath,
			EnumDescriptorProto: ed,
			Index:               i,
			ForcePrefixedName:   r.standalone,
		}
		file.Enums = append(file.Enums, e)
		r.enums[e.FQEN()] = e
		glog.V(1).Infof("register enum name: %s", e.FQEN())
	}
}

// LookupMsg looks up a message type by "name".
// It tries to resolve "name" from "location" if "name" is a relative message name.
func (r *Registry) LookupMsg(location, name string) (*Message, error) {
	glog.V(1).Infof("lookup %s from %s", name, location)
	if strings.HasPrefix(name, ".") {
		m, ok := r.msgs[name]
		if !ok {
			return nil, fmt.Errorf("no message found: %s", name)
		}
		return m, nil
	}

	if !strings.HasPrefix(location, ".") {
		location = fmt.Sprintf(".%s", location)
	}
	components := strings.Split(location, ".")
	for len(components) > 0 {
		fqmn := strings.Join(append(components, name), ".")
		if m, ok := r.msgs[fqmn]; ok {
			return m, nil
		}
		components = components[:len(components)-1]
	}
	return nil, fmt.Errorf("no message found: %s", name)
}

// LookupEnum looks up a enum type by "name".
// It tries to resolve "name" from "location" if "name" is a relative enum name.
func (r *Registry) LookupEnum(location, name string) (*Enum, error) {
	glog.V(1).Infof("lookup enum %s from %s", name, location)
	if strings.HasPrefix(name, ".") {
		e, ok := r.enums[name]
		if !ok {
			return nil, fmt.Errorf("no enum found: %s", name)
		}
		return e, nil
	}

	if !strings.HasPrefix(location, ".") {
		location = fmt.Sprintf(".%s", location)
	}
	components := strings.Split(location, ".")
	for len(components) > 0 {
		fqen := strings.Join(append(components, name), ".")
		if e, ok := r.enums[fqen]; ok {
			return e, nil
		}
		components = components[:len(components)-1]
	}
	return nil, fmt.Errorf("no enum found: %s", name)
}

// LookupFile looks up a file by name.
func (r *Registry) LookupFile(name string) (*File, error) {
	f, ok := r.files[name]
	if !ok {
		return nil, fmt.Errorf("no such file given: %s", name)
	}
	return f, nil
}

// AddPkgMap adds a mapping from a .proto file to proto package name.
func (r *Registry) AddPkgMap(file, protoPkg string) {
	r.pkgMap[file] = protoPkg
}

// SetPrefix registers the prefix to be added to go package paths generated from proto package names.
func (r *Registry) SetPrefix(prefix string) {
	r.prefix = prefix
}

// SetStandalone registers standalone flag to control package prefix
func (r *Registry) SetStandalone(standalone bool) {
	r.standalone = standalone
}

// SetRecursiveDepth records the max recursion count
func (r *Registry) SetRecursiveDepth(count int) {
	r.recursiveDepth = count
}

// GetRecursiveDepth returns the max recursion count
func (r *Registry) GetRecursiveDepth() int {
	return r.recursiveDepth
}

// ReserveGoPackageAlias reserves the unique alias of go package.
// If succeeded, the alias will be never used for other packages in generated go files.
// If failed, the alias is already taken by another package, so you need to use another
// alias for the package in your go files.
func (r *Registry) ReserveGoPackageAlias(alias, pkgpath string) error {
	if taken, ok := r.pkgAliases[alias]; ok {
		if taken == pkgpath {
			return nil
		}
		return fmt.Errorf("package name %s is already taken. Use another alias", alias)
	}
	r.pkgAliases[alias] = pkgpath
	return nil
}

// GetAllFQMNs returns a list of all FQMNs
func (r *Registry) GetAllFQMNs() []string {
	var keys []string
	for k := range r.msgs {
		keys = append(keys, k)
	}
	return keys
}

// GetAllFQENs returns a list of all FQENs
func (r *Registry) GetAllFQENs() []string {
	var keys []string
	for k := range r.enums {
		keys = append(keys, k)
	}
	return keys
}

// SetAllowMerge controls whether generation one OpenAPI file out of multiple protos
func (r *Registry) SetAllowMerge(allow bool) {
	r.allowMerge = allow
}

// IsAllowMerge whether generation one OpenAPI file out of multiple protos
func (r *Registry) IsAllowMerge() bool {
	return r.allowMerge
}

// SetMergeFileName controls the target OpenAPI file name out of multiple protos
func (r *Registry) SetMergeFileName(mergeFileName string) {
	r.mergeFileName = mergeFileName
}

// SetIncludePackageInTags controls whether the package name defined in the `package` directive
// in the proto file can be prepended to the gRPC service name in the `Tags` field of every operation.
func (r *Registry) SetIncludePackageInTags(allow bool) {
	r.includePackageInTags = allow
}

// IsIncludePackageInTags checks whether the package name defined in the `package` directive
// in the proto file can be prepended to the gRPC service name in the `Tags` field of every operation.
func (r *Registry) IsIncludePackageInTags() bool {
	return r.includePackageInTags
}

// SetUseJSONNamesForFields sets useJSONNamesForFields
func (r *Registry) SetUseJSONNamesForFields(use bool) {
	r.useJSONNamesForFields = use
}

// GetUseJSONNamesForFields returns useJSONNamesForFields
func (r *Registry) GetUseJSONNamesForFields() bool {
	return r.useJSONNamesForFields
}

// GetMergeFileName return the target merge OpenAPI file name
func (r *Registry) GetMergeFileName() string {
	return r.mergeFileName
}

// SetUseGoTemplate sets useGoTemplate
func (r *Registry) SetUseGoTemplate(use bool) {
	r.useGoTemplate = use
}

// GetUseGoTemplate returns useGoTemplate
func (r *Registry) GetUseGoTemplate() bool {
	return r.useGoTemplate
}

// SetEnumsAsInts set enumsAsInts
func (r *Registry) SetEnumsAsInts(enumsAsInts bool) {
	r.enumsAsInts = enumsAsInts
}

// GetEnumsAsInts returns enumsAsInts
func (r *Registry) GetEnumsAsInts() bool {
	return r.enumsAsInts
}

// SetOmitEnumDefaultValue sets omitEnumDefaultValue
func (r *Registry) SetOmitEnumDefaultValue(omit bool) {
	r.omitEnumDefaultValue = omit
}

// GetOmitEnumDefaultValue returns omitEnumDefaultValue
func (r *Registry) GetOmitEnumDefaultValue() bool {
	return r.omitEnumDefaultValue
}

// SetVisibilityRestrictionSelectors sets the visibility restriction selectors.
func (r *Registry) SetVisibilityRestrictionSelectors(selectors []string) {
	r.visibilityRestrictionSelectors = make(map[string]bool)
	for _, selector := range selectors {
		r.visibilityRestrictionSelectors[strings.TrimSpace(selector)] = true
	}
}

// GetVisibilityRestrictionSelectors retrieves he visibility restriction selectors.
func (r *Registry) GetVisibilityRestrictionSelectors() map[string]bool {
	return r.visibilityRestrictionSelectors
}

// SetProto3OptionalNullable set proto3OtionalNullable
func (r *Registry) SetProto3OptionalNullable(proto3OtionalNullable bool) {
	r.proto3OptionalNullable = proto3OtionalNullable
}

// GetProto3OptionalNullable returns proto3OtionalNullable
func (r *Registry) GetProto3OptionalNullable() bool {
	return r.proto3OptionalNullable
}

func (r *Registry) FieldName(f *Field) string {
	if r.useJSONNamesForFields {
		return f.GetJsonName()
	}
	return f.GetName()
}
