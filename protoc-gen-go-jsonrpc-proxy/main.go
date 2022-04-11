// Command protoc-gen-grpc-gateway is a plugin for Google protocol buffer
// compiler to generate a reverse-proxy, which converts incoming RESTful
// HTTP/1 requests gRPC invocation.
// You rarely need to run this program directly. Instead, put this program
// into your $PATH with a name "protoc-gen-grpc-gateway" and run
//   protoc --grpc-gateway_out=output_directory path/to/input.proto
//
// See README.md for more details.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang/glog"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/yxlimo/go-jsonrpc-gateway/internal/descriptor"
	"github.com/yxlimo/go-jsonrpc-gateway/protoc-gen-go-jsonrpc-proxy/internal/gengateway"
)

var (
	standalone  = flag.Bool("standalone", false, "generates a standalone gateway package, which imports the target service package")
	versionFlag = flag.Bool("version", false, "print the current version")
)

// Variables set by goreleaser at build time
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if *versionFlag {
		fmt.Printf("Version %v, commit %v, built at %v\n", version, commit, date)
		os.Exit(0)
	}

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		reg := descriptor.NewRegistry()
		err := applyFlags(reg)
		if err != nil {
			return err
		}

		generator := gengateway.New(reg, *standalone)
		glog.V(1).Infof("Parsing code generator request")

		if err := reg.LoadFromPlugin(gen); err != nil {
			return err
		}
		var targets []*descriptor.File
		for _, target := range gen.Request.FileToGenerate {
			f, err := reg.LookupFile(target)
			if err != nil {
				return err
			}
			targets = append(targets, f)
		}
		files, err := generator.Generate(targets)
		if err != nil {
			return err
		}

		for _, f := range files {
			glog.V(1).Infof("NewGeneratedFile %q in %s", f.GetName(), f.GoPkg)
			genFile := gen.NewGeneratedFile(f.GetName(), protogen.GoImportPath(f.GoPkg.Path))
			if _, err := genFile.Write([]byte(f.GetContent())); err != nil {
				return err
			}
		}

		glog.V(1).Info("Processed code generator request")

		return err
	})
}

func parseFlags(reg *descriptor.Registry, parameter string) {
	if parameter == "" {
		return
	}

	for _, p := range strings.Split(parameter, ",") {
		spec := strings.SplitN(p, "=", 2)
		if len(spec) == 1 {
			if err := flag.CommandLine.Set(spec[0], ""); err != nil {
				glog.Fatalf("Cannot set flag %s", p)
			}
			continue
		}

		name, value := spec[0], spec[1]

		if strings.HasPrefix(name, "M") {
			reg.AddPkgMap(name[1:], value)
			continue
		}
		if err := flag.CommandLine.Set(name, value); err != nil {
			glog.Fatalf("Cannot set flag %s", p)
		}
	}
}

func applyFlags(reg *descriptor.Registry) error {
	reg.SetStandalone(*standalone)
	return nil
}
