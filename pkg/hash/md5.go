package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5Hash(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}
