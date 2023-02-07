package dao

import (
	"fmt"
	"testing"
)

func Test_CreateKey(t *testing.T) {
	err := CreateKey(1)
	if err != nil {
		t.Error(err)
	}
}

func Test_FindKeys(t *testing.T) {
	fmt.Println(FindKeys())
}

func Test_KeyBind(t *testing.T) {
	fmt.Println(BindKey("4d07375d-a3c3-11ed-91cf-8045dde2090b", "createbybullet@163.com"))
}

func Test_FindTrafficByUser(t *testing.T) {
	fmt.Println(FindTrafficByUser("createbybullet@163.com"))
}
