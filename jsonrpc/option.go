package jsonrpc

import "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

// ServeMuxOption is an option that can be given to a ServeMux on construction.
type ServeMuxOption func(*ServeMux)

func WithMarshalerOption(marshaler runtime.Marshaler) ServeMuxOption {
	return func(s *ServeMux) {
		s.marshaller = marshaler
	}
}

func WithGatewayRuntimeOptions(opts ...runtime.ServeMuxOption) ServeMuxOption {
	return func(s *ServeMux) {
		for _, opt := range opts {
			opt(s.mux)
		}
	}
}
