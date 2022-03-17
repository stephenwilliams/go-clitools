package hashing

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
