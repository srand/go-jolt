package jolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAttr(t *testing.T) {
	a := NewListAttribute()
	a.Append("a")
	a.Append("b")

	assert.Equal(t, a.Prefix("-D").Join(" ").String(), "-Da -Db")
	assert.Equal(t, a.Suffix("-D").Join(" ").String(), "a-D b-D")
	assert.Equal(t, a.Join("=").String(), "a=b")
	assert.Equal(t, a.String(), "a b")
}
