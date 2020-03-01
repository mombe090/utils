package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMd5(v string) string  {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(v))
	return hex.EncodeToString(hash.Sum(nil))
}
