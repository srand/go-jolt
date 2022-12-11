package jolt

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepfileParser(t *testing.T) {
	file := `
output1: in\ put1 input2 \
  input3 \
   input4

output2: input1  input4

build.dir/proto/dog.pb.cc \
build.dir/proto/dog.pb.h: proto/dog.proto
`

	reader := bytes.NewBufferString(file)
	deps, err := ParseDepfile(reader)
	assert.NoError(t, err)

	inputs, ok := deps["build.dir/proto/dog.pb.cc"]
	assert.True(t, ok)
	if !reflect.DeepEqual(inputs, []string{"proto/dog.proto"}) {
		t.Fatal(inputs)
	}

	inputs, ok = deps["build.dir/proto/dog.pb.h"]
	assert.True(t, ok)
	if !reflect.DeepEqual(inputs, []string{"proto/dog.proto"}) {
		t.Fatal(inputs)
	}

	inputs, ok = deps["output1"]
	assert.True(t, ok)
	if !reflect.DeepEqual(inputs, []string{"in put1", "input2", "input3", "input4"}) {
		t.Fatal(inputs)
	}

	inputs, ok = deps["output2"]
	assert.True(t, ok)
	if !reflect.DeepEqual(inputs, []string{"input1", "input4"}) {
		t.Fatal(inputs)
	}
}
