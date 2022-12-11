package jolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapAttr(t *testing.T) {
	a := NewMapAttribute()
	a.Set("a", "1")
	a.Set("b", "2")

	assert.Equal(t, a.Prefix("-D").String(), "-Da=1 -Db=2")
	assert.Equal(t, a.Suffix("-D").String(), "a=1-D b=2-D")
	assert.Equal(t, a.Join(":").String(), "a:1 b:2")
	assert.Equal(t, a.Join(":").Prefix("-D").String(), "-Da:1 -Db:2")
	assert.Equal(t, a.String(), "a=1 b=2")

	a.Set("c", "")
	assert.Equal(t, a.String(), "a=1 b=2 c")
}
