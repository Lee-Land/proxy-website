package tools

import (
	"crypto/md5"
	"fmt"
)

const securate = "@liu_proxy@"

func PasswordEncode(pwd string) string {
	has := md5.Sum(append([]byte(pwd), []byte(securate)...))
	return fmt.Sprintf("%x", has)
}
