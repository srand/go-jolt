package jolt

type Influence interface {
	Digest() (Digest, error)
}
