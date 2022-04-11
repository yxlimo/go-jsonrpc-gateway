
.PHONY: install
install:
	go install ./protoc-gen-go-jsonrpc-proxy

.PHONY: gen-pb
gen-pb: install
	buf generate