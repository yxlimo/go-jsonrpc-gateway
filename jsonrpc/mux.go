package jsonrpc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// ServerCodec implements reading, parsing and writing RPC messages for the server side of
// a RPC session. Implementations must be go-routine safe since the codec can be called in
// multiple go-routines concurrently.
type ServerCodec interface {
	readBatch() (msgs []*jsonrpcMessage, isBatch bool, err error)
	close()

	jsonWriter
}

// jsonWriter can write JSON messages to its underlying connection.
// Implementations must be safe for concurrent use.
type jsonWriter interface {
	writeJSON(context.Context, interface{}) error
	// Closed returns a channel which is closed when the connection is closed.
	closed() <-chan interface{}
}

type HandleFunc func(req *http.Request, marshaller runtime.Marshaler, rawBody json.RawMessage) (json.RawMessage, context.Context, error)

type ServeMux struct {
	mux        *runtime.ServeMux
	marshaller runtime.Marshaler
	handlers   map[string]HandleFunc
}

func NewServeMux(opts ...ServeMuxOption) *ServeMux {
	mux := &ServeMux{
		mux: runtime.NewServeMux(),
		marshaller: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
		handlers: make(map[string]HandleFunc),
	}
	for _, opt := range opts {
		opt(mux)
	}

	return mux
}

func (s *ServeMux) RuntimeMux() *runtime.ServeMux {
	return s.mux
}

func (s *ServeMux) Register(method string, handler HandleFunc) {
	s.handlers[method] = handler
}

func (s *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = s.mux.HandlePath("GET", "/**", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.Error(w, "GET method not allowed", http.StatusMethodNotAllowed)
	})
	_ = s.mux.HandlePath("PUT", "/**", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.Error(w, "PUT method not allowed", http.StatusMethodNotAllowed)
	})
	_ = s.mux.HandlePath("DELETE", "/**", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.Error(w, "DELETE method not allowed", http.StatusMethodNotAllowed)
	})
	_ = s.mux.HandlePath("POST", "/**", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		if status, err := ValidateRequest(r); err != nil {
			http.Error(w, err.Error(), status)
			return
		}
		codec := NewHTTPServerConn(r, w, s.marshaller)
		msg, isBatch, err := codec.readBatch()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if isBatch {
			http.Error(w, "batch request not supported", http.StatusBadRequest)
			return
		}
		h, ok := s.handlers[msg[0].Method]
		if !ok {
			httpErrorHandler(r.Context(), s, s.marshaller, w, msg[0], status.New(codes.Unimplemented, "method not implemented").Err())
			return
		}
		resp, newCtx, err := h(r, s.marshaller, msg[0].Params)
		if err != nil {
			httpErrorHandler(newCtx, s, s.marshaller, w, msg[0], err)
			return
		}
		byes, err := s.marshaller.Marshal(resp)
		if err != nil {
			httpErrorHandler(newCtx, s, s.marshaller, w, msg[0], err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = codec.writeJSON(newCtx, &jsonrpcMessage{
			Version: "2.0",
			ID:      msg[0].ID,
			Method:  msg[0].Method,
			Result:  byes,
		})
	})

	s.mux.ServeHTTP(w, r)
}
