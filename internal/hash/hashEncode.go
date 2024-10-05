package hashEncode

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHashSHA(longUrl string) string {
	hash := sha256.Sum256([]byte(longUrl))
	return hex.EncodeToString(hash[:])[:10]
}
