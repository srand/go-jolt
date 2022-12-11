package jolt

import (
	"testing"
)

func TestStringInfluence(t *testing.T) {
	fi := StringInfluence("testing")
	digest, err := fi.Digest()
	if err != nil {
		t.Error(err)
	}
	if digest != Digest("dc724af18fbdd4e59189f5fe768a5f8311527050") {
		t.Error(digest)
	}
}
