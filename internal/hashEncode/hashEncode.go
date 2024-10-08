package hashEncode

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHashSHA256(longUrl string, numberCaracteres int) string {
	hash := sha256.Sum256([]byte(longUrl))
	return hex.EncodeToString(hash[:])[:numberCaracteres]
}

func GenerateHashMD5(longUrl string, numberCaracteres int) string {
	hash := md5.Sum([]byte(longUrl))
	shortUrl := hex.EncodeToString(hash[:])[:8]
	return shortUrl
}
