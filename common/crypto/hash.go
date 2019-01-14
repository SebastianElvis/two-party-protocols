package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateHash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
