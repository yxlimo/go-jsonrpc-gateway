package jsonrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/assert"
)

func TestMuxServeHTTP(t *testing.T) {
	for i, spec := range []struct {
		reqMethod  string
		reqPath    string
		reqContent map[string]interface{}
		headers    map[string]string
		jrpcMethod string

		respStatus  int
		respContent map[string]interface{}
	}{
		{
			reqMethod:  "GET",
			reqPath:    "/",
			respStatus: http.StatusMethodNotAllowed,
		},
		{
			reqMethod:  "GET",
			reqPath:    "/aaaa",
			respStatus: http.StatusMethodNotAllowed,
		},
		{
			reqMethod:  "PUT",
			reqPath:    "/",
			respStatus: http.StatusMethodNotAllowed,
		},
		{
			reqMethod:  "DELETE",
			reqPath:    "/",
			respStatus: http.StatusMethodNotAllowed,
		},
		{
			reqMethod:  "POST",
			reqPath:    "/",
			jrpcMethod: "Service.Hello",
			reqContent: map[string]interface{}{
				"jsonrpc": "2.0",
			},
			headers: map[string]string{
				"Content-Type": "application/no-json",
			},
			respStatus: http.StatusUnsupportedMediaType,
		},
		{
			reqMethod:  "POST",
			reqPath:    "/",
			jrpcMethod: "Service.Hello",
			reqContent: map[string]interface{}{
				"jsonrpc": "2.0",
				"method":  "Service.Greet",
				"id":      "1",
				"params": map[string]string{
					"name": "world",
				},
			},
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			respStatus: http.StatusNotImplemented,
		},
		{
			reqMethod:  "POST",
			reqPath:    "/",
			jrpcMethod: "Service.Hello",
			reqContent: map[string]interface{}{
				"jsonrpc": "2.0",
				"method":  "Service.Hello",
				"id":      "1",
				"params": map[string]string{
					"name": "world",
				},
			},
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			respStatus: http.StatusOK,
			respContent: map[string]interface{}{
				"jsonrpc": "2.0",
				"method":  "Service.Hello",
				"id":      "1",
				"result": map[string]interface{}{
					"name": "world",
				},
			},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			mux := NewServeMux()
			func(jrpcMethod string) {
				mux.Register(jrpcMethod, func(req *http.Request, marshaller runtime.Marshaler, rawBody json.RawMessage) (json.RawMessage, context.Context, error) {
					return rawBody, req.Context(), nil
				})
			}(spec.jrpcMethod)

			reqUrl := fmt.Sprintf("https://host.example%s", spec.reqPath)
			reqReader := bytes.NewReader(nil)
			if spec.reqContent != nil {
				raw, _ := json.Marshal(spec.reqContent)
				reqReader = bytes.NewReader(raw)
			}
			r, err := http.NewRequest(spec.reqMethod, reqUrl, reqReader)
			if err != nil {
				t.Fatalf("http.NewRequest(%q, %q, nil) failed with %v; want success", spec.reqMethod, reqUrl, err)
			}
			for name, value := range spec.headers {
				r.Header.Set(name, value)
			}
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)

			if got, want := w.Code, spec.respStatus; got != want {
				t.Errorf("w.Code = %d; want %d; patterns=%v; req=%v", got, want, spec.reqPath, r)
			}
			if spec.respContent != nil {
				jsonResp := map[string]interface{}{}
				_ = json.NewDecoder(w.Body).Decode(&jsonResp)
				assert.Equal(t, spec.respContent, jsonResp)
			}

		})
	}
}
