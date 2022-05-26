package client

import (
	"strings"
	"testing"
)

func Test_generateFallbackId(t *testing.T) {

	got := generateFallBackTraceId()

	if !strings.HasPrefix(got, "fallback-") {
		t.Fatalf("failed to get fallback id Got: %s, Expected : %s...", got, "fallback-")
	}

}
