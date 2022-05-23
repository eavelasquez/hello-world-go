package crypto

import "crypto/sha256"

func NewSHA256(data []byte) ([]byte, error) {
	h := sha256.Sum256(data)
	return h[:], nil
}
