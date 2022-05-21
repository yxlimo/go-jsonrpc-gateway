package openapi

import (
	"encoding/json"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

var _ pgs.Module = &Openapi{}

type Openapi struct {
	base *pgs.ModuleBase
	ctx  pgsgo.Context

	schemas        map[string]*openapiSchemaObject
	paths          map[string]*openapiPathObject
	importMessages map[string]pgs.Message
	msgTypes       map[string]string
}

func New() *Openapi {
	return &Openapi{
		base:           &pgs.ModuleBase{},
		schemas:        make(map[string]*openapiSchemaObject),
		paths:          make(map[string]*openapiPathObject),
		msgTypes:       make(map[string]string),
		importMessages: make(map[string]pgs.Message),
	}
}

func (o *Openapi) Name() string {
	return "jsonrpc-openapiv3"
}

func (o *Openapi) InitContext(c pgs.BuildContext) {
	o.base.InitContext(c)
	o.ctx = pgsgo.InitContext(c.Parameters())
}

func (o *Openapi) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		o.generate(t)
	}
	return o.base.Artifacts()
}

func (s *Openapi) generate(file pgs.File) {
	s.registerImports(file.Imports())
	s.registerMessageTypes(file.AllMessages())
	for _, message := range file.AllMessages() {
		s.schemas[s.messageRefName(message.FullyQualifiedName())] = s.genMessage(message)
	}
	for _, service := range file.Services() {
		for _, method := range service.Methods() {
			s.paths["/"+method.Name().LowerSnakeCase().String()] = s.genMethod(method)
		}
	}
	object := openapiObject{
		Version: "3.0.0",
		Info: openapiInfoObject{
			Title:   file.Name().String(),
			Version: "0.0.1",
		},
	}
	object.Paths = s.paths
	object.Components.Schemas = s.schemas
	content, _ := json.Marshal(&object)
	name := s.ctx.OutputPath(file).SetExt(".openapi.json")
	s.base.OverwriteCustomFile(name.String(), string(content), 0644)
}

func (s *Openapi) registerImports(imports []pgs.File) {
	for _, file := range imports {
		s.base.Debugf("registering imports: %s", file.InputPath().String())
		for _, msg := range file.AllMessages() {
			s.base.Debugf("registering import message: %s", msg.FullyQualifiedName())
			s.importMessages[msg.FullyQualifiedName()] = msg
			s.msgTypes[msg.FullyQualifiedName()] = msg.FullyQualifiedName()
		}
	}
}

func (s *Openapi) registerMessageTypes(messages []pgs.Message) {
	for _, msg := range messages {
		for _, entry := range msg.MapEntries() {
			s.base.Debugf("registering map entries: %s", entry.FullyQualifiedName())
			valueField := entry.Fields()[1]
			if valueField.Type().ProtoType() == pgs.MessageT {
				s.msgTypes[entry.FullyQualifiedName()] = valueField.Descriptor().GetTypeName()
			}
		}
		s.msgTypes[msg.FullyQualifiedName()] = msg.FullyQualifiedName()
	}
}

func (s *Openapi) Parameters() pgs.Parameters {
	return s.base.Parameters()
}

func (s *Openapi) genMethod(m pgs.Method) *openapiPathObject {
	return &openapiPathObject{
		Post: &openapiOperationObject{
			Summary:     m.SourceCodeInfo().LeadingComments(),
			RequestBody: s.jsonrpcRequestSchema(m.Name().UpperCamelCase().String(), s.genSchemaFromMsg(m.Input())),
			Responses: map[string]*openapiResponseObject{
				"200": s.jsonrpcResponseSchema(m.Name().UpperCamelCase().String(), s.genSchemaFromMsg(m.Output())),
			},
		},
	}
}

func (s *Openapi) genSchemaFromMsg(msg pgs.Message) *openapiSchemaObject {
	if wkt, ok := wktSchemas[msg.FullyQualifiedName()]; ok {
		return wkt
	}
	s.genImportSchemaIfExist(msg.FullyQualifiedName())
	return &openapiSchemaObject{
		Ref: "#/components/schemas/" + s.messageRefName(msg.FullyQualifiedName()),
	}
}

func (s *Openapi) messageRefName(fqn string) string {
	packageNames := strings.Split(fqn, ".")
	return packageNames[len(packageNames)-2] + "." + packageNames[len(packageNames)-1]
}

func (s *Openapi) genMessage(msg pgs.Message) *openapiSchemaObject {
	schema := &openapiSchemaObject{Type: "object", Properties: make(map[string]*openapiSchemaObject, len(msg.Fields()))}
	for _, field := range msg.Fields() {
		schema.Properties[field.Name().LowerSnakeCase().String()] = s.genSchemaFromField(field)
	}
	return schema
}

func (s *Openapi) genSchemaFromField(field pgs.Field) *openapiSchemaObject {
	if field.Type().IsEnum() {
		enums := field.Type().Enum()
		enumsStr := make([]string, 0, len(enums.Values()))
		for _, v := range enums.Values() {
			enumsStr = append(enumsStr, v.Name().String())
		}
		return &openapiSchemaObject{
			Type: "string",
			Enum: enumsStr,
		}
	}
	if field.Type().IsRepeated() {
		return &openapiSchemaObject{
			Type:  "array",
			Items: s.genSchemaFromFieldProtoType(field, field.Type().Element().ProtoType()),
		}
	}
	if field.Type().IsMap() {
		if field.Type().Key().ProtoType() == pgs.StringT {
			return &openapiSchemaObject{
				Type:                 "object",
				AdditionalProperties: s.genSchemaFromFieldProtoType(field, field.Type().Element().ProtoType()),
			}
		}
	}
	return s.genSchemaFromFieldProtoType(field, field.Type().ProtoType())
}

func (s *Openapi) genSchemaFromFieldProtoType(field pgs.Field, typ pgs.ProtoType) *openapiSchemaObject {
	switch typ {
	case pgs.DoubleT:
		return &openapiSchemaObject{Type: "string", Format: "double"}
	case pgs.FloatT:
		return &openapiSchemaObject{Type: "number", Format: "float"}
	case pgs.Int64T, pgs.SFixed64, pgs.SInt64:
		return &openapiSchemaObject{Type: "string", Format: "int64"}
	case pgs.UInt64T, pgs.Fixed64T:
		return &openapiSchemaObject{Type: "string", Format: "uint64"}
	case pgs.Int32T, pgs.SFixed32, pgs.SInt32:
		return &openapiSchemaObject{Type: "integer", Format: "int32"}
	case pgs.UInt32T, pgs.Fixed32T:
		return &openapiSchemaObject{Type: "integer", Format: "int64"}
	case pgs.BoolT:
		return &openapiSchemaObject{Type: "boolean"}
	case pgs.StringT:
		return &openapiSchemaObject{Type: "string"}
	case pgs.MessageT:
		if wktSchema, ok := wktSchemas[field.Descriptor().GetTypeName()]; ok {
			return wktSchema
		}
		msgTypeName := s.msgTypes[field.Descriptor().GetTypeName()]
		s.genImportSchemaIfExist(msgTypeName)
		s.base.Debugf("generate reference schema: %s", field.Descriptor().GetTypeName())
		return &openapiSchemaObject{
			Ref: "#/components/schemas/" + s.messageRefName(msgTypeName),
		}

	case pgs.BytesT:
		return &openapiSchemaObject{Type: "string", Format: "byte"}
	}
	return &openapiSchemaObject{Type: "string"}
}

func (s *Openapi) genImportSchemaIfExist(fqn string) {
	if msg, exist := s.importMessages[fqn]; exist {
		s.schemas[s.messageRefName(msg.FullyQualifiedName())] = s.genMessage(msg)
	}
}
