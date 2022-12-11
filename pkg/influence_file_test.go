package jolt

import (
	"log"
	"os"
	"testing"
)

func TestFileInfluence(t *testing.T) {
	file, err := os.CreateTemp("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	file.WriteString("testing")
	file.Sync()

	fi := FileInfluence(file.Name())
	digest, err := fi.Digest()
	if err != nil {
		t.Error(err)
	}
	if digest != Digest("dc724af18fbdd4e59189f5fe768a5f8311527050") {
		t.Error(digest)
	}
}
