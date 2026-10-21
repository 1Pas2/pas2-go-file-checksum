package pas2

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"io"
)

func Digest(reader io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
func Verify(reader io.Reader, expected string) (bool, error) {
	raw, err := hex.DecodeString(expected)
	if err != nil || len(raw) != sha256.Size {
		return false, fmt.Errorf("expected 64 hex characters")
	}
	actual, err := Digest(reader)
	if err != nil {
		return false, err
	}
	got, _ := hex.DecodeString(actual)
	return subtle.ConstantTimeCompare(got, raw) == 1, nil
}
