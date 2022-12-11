package jolt

import "os"

type fileInfluence struct {
	path string
}

func FileInfluence(path string) Influence {
	return &fileInfluence{path: path}
}

func (fi *fileInfluence) Digest() (Digest, error) {
	fd, err := os.Open(fi.path)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	return Sha1(fd)
}
