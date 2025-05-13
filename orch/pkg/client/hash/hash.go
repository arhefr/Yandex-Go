package hash

import (
	"crypto/sha1"
	"fmt"
)

type PasswordHasher interface {
	Hash(password string) string
}

type Hasher struct {
	salt string
}

func NewHasher(salt string) *Hasher {
	return &Hasher{salt: salt}
}

func (h *Hasher) Hash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt)))
}
