package tools

import (
	"fmt"
	"testing"
)

func Test_PasswordEncode(t *testing.T) {
	pwd := "123456"
	s := PasswordEncode(pwd)
	fmt.Println(s)
}
