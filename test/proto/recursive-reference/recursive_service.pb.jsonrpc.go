// Code generated by protoc-gen-jsonrpc-gateway. DO NOT EDIT.
// source: test/proto/recursive-reference/recursive_service.proto

// Package recursive_reference is a reverse proxy.

// It translates gRPC into JSON-RPC APIs.
package recursive_reference

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/yxlimo/go-jsonrpc-gateway/jsonrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = metadata.Join
var _ = json.Marshal
var _ = jsonrpc.NewHTTPServerConn

func request_Recursive_RecursiveCall_jsonrpc(ctx context.Context, marshaler runtime.Marshaler, client RecursiveClient, raw json.RawMessage) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FooRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(bytes.NewReader(raw)).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.RecursiveCall(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

// RegisterRecursiveJSONRPCHandlerFromEndpoint is same as RegisterRecursiveJSONRPCHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterRecursiveJSONRPCHandlerFromEndpoint(ctx context.Context, mux *jsonrpc.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterRecursiveJSONRPCHandler(ctx, mux, conn)
}

// RegisterRecursiveJSONRPCHandler registers the http handlers for service Recursive to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterRecursiveJSONRPCHandler(ctx context.Context, mux *jsonrpc.ServeMux, conn *grpc.ClientConn) error {
	return RegisterRecursiveJSONRPCHandlerClient(ctx, mux, NewRecursiveClient(conn))
}

// RegisterRecursiveJSONRPCHandlerClient registers the http handlers for service Recursive
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "RecursiveClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "RecursiveClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "RecursiveClient" to call the correct interceptors.
func RegisterRecursiveJSONRPCHandlerClient(ctx context.Context, mux *jsonrpc.ServeMux, client RecursiveClient) error {

	mux.Register("RecursiveCall", func(req *http.Request, marshaller runtime.Marshaler, rawBody json.RawMessage) (json.RawMessage, context.Context, error) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux.RuntimeMux(), req, "/jsonrpc.gateway.test.proto.recursive_reference.Recursive/RecursiveCall")
		if err != nil {
			return nil, ctx, err
		}
		resp, md, err := request_Recursive_RecursiveCall_jsonrpc(ctx, marshaller, client, rawBody)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			return nil, ctx, err
		}
		rawResp, err := marshaller.Marshal(resp)
		if err != nil {
			return nil, ctx, err
		}
		return rawResp, ctx, nil

	})

	return nil
}
