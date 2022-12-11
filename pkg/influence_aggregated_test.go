package jolt

import (
	"fmt"
	"testing"
)

type UndefinedInfluence struct{}

func (*UndefinedInfluence) Digest() (Digest, error) {
	return "", fmt.Errorf("Undefined")
}

func TestAggregatedInfluenceDigest(t *testing.T) {
	ai := NewAggregatedInfluence()

	ai.AddInfluence(StringInfluence("testing"))
	ai.AddInfluence(StringInfluence("testing"))

	digest, err := ai.Digest()
	if err != nil {
		t.Error(err)
	}
	if digest != Digest("cf49a2f6e6036c28d41ee8f392a6802f186ae7f7") {
		t.Error(digest)
	}
}

func TestAggregatedInfluenceError(t *testing.T) {
	ai := NewAggregatedInfluence()

	ai.AddInfluence(StringInfluence("testing"))
	ai.AddInfluence(&UndefinedInfluence{})

	digest, err := ai.Digest()
	if err == nil {
		t.Error("Didn't fail", digest)
	}
}
