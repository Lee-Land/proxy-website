package tools

import "testing"

func Test_JWT(t *testing.T) {
	username := "createbybullet@163.com"
	s, err := GenerateToken(username)
	if err != nil {
		t.Error(err)
		return
	}
	t2, err2 := ParseToken(s)
	if err2 != nil {
		t.Error(err2)
		return
	}
	if !t2.Valid {
		t.Error("生成的token【" + s + "】无法认证")
	}
}
