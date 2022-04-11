package gengateway

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/yxlimo/go-jsonrpc-gateway/internal/descriptor"
)

func crossLinkFixture(f *descriptor.File) *descriptor.File {
	for _, m := range f.Messages {
		m.File = f
	}
	for _, svc := range f.Services {
		svc.File = f
		for _, m := range svc.Methods {
			m.Service = svc
		}
	}
	return f
}

func TestApplyTemplateHeader(t *testing.T) {
	msgdesc := &descriptorpb.DescriptorProto{
		Name: proto.String("ExampleMessage"),
	}
	meth := &descriptorpb.MethodDescriptorProto{
		Name:       proto.String("Example"),
		InputType:  proto.String("ExampleMessage"),
		OutputType: proto.String("ExampleMessage"),
	}
	svc := &descriptorpb.ServiceDescriptorProto{
		Name:   proto.String("ExampleService"),
		Method: []*descriptorpb.MethodDescriptorProto{meth},
	}
	msg := &descriptor.Message{
		DescriptorProto: msgdesc,
	}
	file := descriptor.File{
		FileDescriptorProto: &descriptorpb.FileDescriptorProto{
			Name:        proto.String("example.proto"),
			Package:     proto.String("example"),
			Dependency:  []string{"a.example/b/c.proto", "a.example/d/e.proto"},
			MessageType: []*descriptorpb.DescriptorProto{msgdesc},
			Service:     []*descriptorpb.ServiceDescriptorProto{svc},
		},
		GoPkg: descriptor.GoPackage{
			Path: "example.com/path/to/example/example.pb",
			Name: "example_pb",
		},
		Messages: []*descriptor.Message{msg},
		Services: []*descriptor.Service{
			{
				ServiceDescriptorProto: svc,
				Methods: []*descriptor.Method{
					{
						MethodDescriptorProto: meth,
						RequestType:           msg,
						ResponseType:          msg,
					},
				},
			},
		},
	}
	got, err := applyTemplate(param{File: crossLinkFixture(&file)}, descriptor.NewRegistry())
	if err != nil {
		t.Errorf("applyTemplate(%#v) failed with %v; want success", file, err)
		return
	}
	if want := "package example_pb\n"; !strings.Contains(got, want) {
		t.Errorf("applyTemplate(%#v) = %s; want to contain %s", file, got, want)
	}
}

func TestIdentifierCapitalization(t *testing.T) {
	msgdesc1 := &descriptorpb.DescriptorProto{
		Name: proto.String("Exam_pleRequest"),
	}
	msgdesc2 := &descriptorpb.DescriptorProto{
		Name: proto.String("example_response"),
	}
	meth1 := &descriptorpb.MethodDescriptorProto{
		Name:       proto.String("ExampleGe2t"),
		InputType:  proto.String("Exam_pleRequest"),
		OutputType: proto.String("example_response"),
	}
	meth2 := &descriptorpb.MethodDescriptorProto{
		Name:       proto.String("Exampl_ePost"),
		InputType:  proto.String("Exam_pleRequest"),
		OutputType: proto.String("example_response"),
	}
	svc := &descriptorpb.ServiceDescriptorProto{
		Name:   proto.String("Example"),
		Method: []*descriptorpb.MethodDescriptorProto{meth1, meth2},
	}
	msg1 := &descriptor.Message{
		DescriptorProto: msgdesc1,
	}
	msg2 := &descriptor.Message{
		DescriptorProto: msgdesc2,
	}
	file := descriptor.File{
		FileDescriptorProto: &descriptorpb.FileDescriptorProto{
			Name:        proto.String("example.proto"),
			Package:     proto.String("example"),
			Dependency:  []string{"a.example/b/c.proto", "a.example/d/e.proto"},
			MessageType: []*descriptorpb.DescriptorProto{msgdesc1, msgdesc2},
			Service:     []*descriptorpb.ServiceDescriptorProto{svc},
		},
		GoPkg: descriptor.GoPackage{
			Path: "example.com/path/to/example/example.pb",
			Name: "example_pb",
		},
		Messages: []*descriptor.Message{msg1, msg2},
		Services: []*descriptor.Service{
			{
				ServiceDescriptorProto: svc,
				Methods: []*descriptor.Method{
					{
						MethodDescriptorProto: meth1,
						RequestType:           msg1,
						ResponseType:          msg1,
					},
				},
			},
			{
				ServiceDescriptorProto: svc,
				Methods: []*descriptor.Method{
					{
						MethodDescriptorProto: meth2,
						RequestType:           msg2,
						ResponseType:          msg2,
					},
				},
			},
		},
	}

	got, err := applyTemplate(param{File: crossLinkFixture(&file)}, descriptor.NewRegistry())
	if err != nil {
		t.Errorf("applyTemplate(%#v) failed with %v; want success", file, err)
		return
	}
	if want := `msg, err := client.ExampleGe2T(ctx, &protoReq, grpc.Header(&metadata.HeaderMD)`; !strings.Contains(got, want) {
		t.Errorf("applyTemplate(%#v) = %s; want to contain %s", file, got, want)
	}
	if want := `msg, err := client.ExamplEPost(ctx, &protoReq, grpc.Header(&metadata.HeaderMD)`; !strings.Contains(got, want) {
		t.Errorf("applyTemplate(%#v) = %s; want to contain %s", file, got, want)
	}
	if want := `var protoReq ExamPleRequest`; !strings.Contains(got, want) {
		t.Errorf("applyTemplate(%#v) = %s; want to contain %s", file, got, want)
	}
	if want := `var protoReq ExampleResponse`; !strings.Contains(got, want) {
		t.Errorf("applyTemplate(%#v) = %s; want to contain %s", file, got, want)
	}
}
