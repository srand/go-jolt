package jolt

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

type Digest string

func Sha1(r io.Reader) (Digest, error) {
	hash := sha1.New()
	_, err := io.Copy(hash, r)
	if err != nil {
		return "", err
	}

	digest := hash.Sum(nil)
	hexdigest := hex.EncodeToString(digest)
	return Digest(hexdigest), nil
}
