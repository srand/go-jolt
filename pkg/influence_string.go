package jolt

import "bytes"

type stringInfluence struct {
	data []byte
}

func StringInfluence(data string) Influence {
	return &stringInfluence{data: []byte(data)}
}

func (si *stringInfluence) Digest() (Digest, error) {
	return Sha1(bytes.NewReader(si.data))
}
