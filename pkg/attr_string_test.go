package jolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringAttr(t *testing.T) {
	a := NewStringAttribute("string")
	assert.Equal(t, a.Prefix("-D").String(), "-Dstring")
	assert.Equal(t, a.Suffix("-D").String(), "string-D")
	assert.Equal(t, a.Join(" ").String(), "string")
}
