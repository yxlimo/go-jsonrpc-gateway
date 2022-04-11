package gengateway

import (
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"path"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/yxlimo/go-jsonrpc-gateway/internal/descriptor"
)

var (
	errNoTargetService = errors.New("no target service defined in the file")
)

type generator struct {
	reg         *descriptor.Registry
	baseImports []descriptor.GoPackage
	standalone  bool
}

// New returns a new generator which generates grpc gateway files.
func New(reg *descriptor.Registry, standalone bool) *generator {
	var imports []descriptor.GoPackage
	for _, pkgpath := range []string{
		"bytes",
		"context",
		"encoding/json",
		"io",
		"net/http",
		"github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
		"github.com/yxlimo/go-jsonrpc-gateway/jsonrpc",
		"google.golang.org/protobuf/proto",
		"google.golang.org/grpc",
		"google.golang.org/grpc/codes",
		"google.golang.org/grpc/grpclog",
		"google.golang.org/grpc/metadata",
		"google.golang.org/grpc/status",
	} {
		pkg := descriptor.GoPackage{
			Path: pkgpath,
			Name: path.Base(pkgpath),
		}
		if err := reg.ReserveGoPackageAlias(pkg.Name, pkg.Path); err != nil {
			for i := 0; ; i++ {
				alias := fmt.Sprintf("%s_%d", pkg.Name, i)
				if err := reg.ReserveGoPackageAlias(alias, pkg.Path); err != nil {
					continue
				}
				pkg.Alias = alias
				break
			}
		}
		imports = append(imports, pkg)
	}

	return &generator{
		reg:         reg,
		baseImports: imports,
		standalone:  standalone,
	}
}

func (g *generator) Generate(targets []*descriptor.File) ([]*descriptor.ResponseFile, error) {
	var files []*descriptor.ResponseFile
	for _, file := range targets {
		glog.V(1).Infof("Processing %s", file.GetName())
		code, err := g.generate(file)
		if err == errNoTargetService {
			glog.V(1).Infof("%s: %v", file.GetName(), err)
			continue
		}
		if err != nil {
			return nil, err
		}
		formatted, err := format.Source([]byte(code))
		if err != nil {
			glog.Errorf("format code failed: %v", err)
			_ = ioutil.WriteFile(file.GeneratedFilenamePrefix+".pb.jsonrpc.go", []byte(code), 0644)
			return nil, err
		}
		files = append(files, &descriptor.ResponseFile{
			GoPkg: file.GoPkg,
			CodeGeneratorResponse_File: &pluginpb.CodeGeneratorResponse_File{
				Name:    proto.String(file.GeneratedFilenamePrefix + ".pb.jsonrpc.go"),
				Content: proto.String(string(formatted)),
			},
		})
	}
	return files, nil
}

func (g *generator) generate(file *descriptor.File) (string, error) {
	pkgSeen := make(map[string]bool)
	var imports []descriptor.GoPackage
	for _, pkg := range g.baseImports {
		pkgSeen[pkg.Path] = true
		imports = append(imports, pkg)
	}

	if g.standalone {
		imports = append(imports, file.GoPkg)
	}

	for _, svc := range file.Services {
		for _, m := range svc.Methods {
			pkg := m.RequestType.File.GoPkg
			if pkg == file.GoPkg || pkgSeen[pkg.Path] {
				continue
			}
			pkgSeen[pkg.Path] = true
			imports = append(imports, pkg)
		}
	}
	params := param{
		File:    file,
		Imports: imports,
	}
	return applyTemplate(params, g.reg)
}
