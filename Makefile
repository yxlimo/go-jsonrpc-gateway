
.PHONY: install
install:
	go install ./protoc-gen-go-jsonrpc-proxy
	go install ./protoc-gen-jsonrpc-openapiv3

.PHONY: gen-pb
gen-pb: install
	DEBUG=true buf generate