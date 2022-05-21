package main

import (
	pgs "github.com/lyft/protoc-gen-star"

	"github.com/yxlimo/go-jsonrpc-gateway/protoc-gen-jsonrpc-openapiv3/internal/openapi"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(openapi.New()).Render()
}
