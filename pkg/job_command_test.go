package jolt

import (
	"testing"
)

func TestCommandJobDigest(t *testing.T) {
	cmd := NewCommandJob("echo test")
	digest, err := cmd.Digest()
	if err != nil {
		t.Error(err)
	}
	if digest != Digest("b4adaa61d0acd529003718a7def93308c5412bd6") {
		t.Error("Digest mismatch", digest)
	}
}
