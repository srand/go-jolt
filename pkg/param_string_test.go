package jolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringParam_AcceptedValues(t *testing.T) {
	p := NewStringParam("param", "", []string{"a", "b"}, "")
	assert.NoError(t, p.Set("a"))
	assert.NoError(t, p.Set("b"))
	assert.Error(t, p.Set("c"))

	p = NewStringParam("param", "", []string{}, "")
	assert.NoError(t, p.Set("a"))
	assert.NoError(t, p.Set("b"))
	assert.NoError(t, p.Set("c"))

	p = NewStringParam("param", "c", []string{"a", "b"}, "")
	assert.Nil(t, p)
}

func TestStringParam_IsSet(t *testing.T) {
	p := NewStringParam("param", "", []string{"a", "b"}, "")
	assert.False(t, p.IsSet())

	p = NewStringParam("param", "b", []string{"a", "b"}, "")
	assert.True(t, p.IsSet())
}

func TestStringParam_Get(t *testing.T) {
	p := NewStringParam("param", "b", []string{"a", "b"}, "")
	assert.Equal(t, p.Get(), "b")
	assert.NoError(t, p.Set("a"))
	assert.Equal(t, p.Get(), "a")
}
