package ports

import (
	"crypto/md5"
	"fmt"
)

func HashUrl(url string) string {
	hash := md5.Sum([]byte(url))
	return fmt.Sprintf("%x", hash)[:5]
}
