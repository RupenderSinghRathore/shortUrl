package models

import (
	"crypto/md5"
	"encoding/hex"
)

func CreateHash(url string) string {
	hash := md5.Sum([]byte(url))
	sliceHash := hash[:]
	return hex.EncodeToString(sliceHash)[:5]
}
