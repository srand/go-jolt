package jolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolParam_AcceptedValues(t *testing.T) {
	p := NewBoolParam("param", false, "")
	assert.NoError(t, p.Set("true"))
	assert.NoError(t, p.Set("yes"))
	assert.NoError(t, p.Set("false"))
	assert.NoError(t, p.Set("no"))
	assert.NoError(t, p.Set(true))
	assert.NoError(t, p.Set(false))
	assert.Error(t, p.Set("asf"))
}

func TestBoolParam_Get(t *testing.T) {
	p := NewBoolParam("param", false, "")
	assert.Equal(t, p.Get(), false)
	assert.NoError(t, p.Set(true))
	assert.Equal(t, p.Get(), true)

	p = NewBoolParam("param", true, "")
	assert.Equal(t, p.Get(), true)
	assert.NoError(t, p.Set(false))
	assert.Equal(t, p.Get(), false)
}
