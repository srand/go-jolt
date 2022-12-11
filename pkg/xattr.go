package jolt

import "golang.org/x/sys/unix"

const (
	XattrJobDigest = "user.jolt.digest"
)

// FIXME: EINTR
func GetXattrDigest(path, key string) (Digest, error) {
	data := make([]byte, 40)
	_, err := unix.Getxattr(path, key, data)
	if err != nil {
		return "", err
	}
	return Digest(data), nil
}

func SetXattrDigest(path, key string, digest Digest) error {
	data := []byte(string(digest))
	err := unix.Setxattr(path, key, data, 0)
	if err != nil {
		return err
	}
	return nil
}
