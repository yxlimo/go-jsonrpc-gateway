package main

import (
	"testing"

	"github.com/yxlimo/go-jsonrpc-gateway/internal/descriptor"
)

func TestParseFlagsEmptyNoPanic(t *testing.T) {
	reg := descriptor.NewRegistry()
	parseFlags(reg, "")
}

func TestParseFlags(t *testing.T) {
	reg := descriptor.NewRegistry()
	parseFlags(reg, "standalone=true")
	if *standalone != true {
		t.Errorf("flag standalone was not set correctly, wanted true got %v", *standalone)
	}
}
