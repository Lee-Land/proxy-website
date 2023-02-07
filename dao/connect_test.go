package dao

import (
	"fmt"
	"testing"
)

func Test_connect(t *testing.T) {
	if DB != nil {
		fmt.Println("数据库连接成功")
	} else {
		fmt.Println("数据库连接失败")
	}

}
