package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashToSHA256(toBeHash string) string {
	hash := sha256.New()
	hash.Write([]byte(toBeHash))

	hashed256 := hex.EncodeToString(hash.Sum(nil))

	return hashed256
}
