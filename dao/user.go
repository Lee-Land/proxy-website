package dao

import (
	"proxy-website/tools"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key"`
	UserName  string    `gorm:"column:username"` //用户名
	Pwd       string    //密码
	CreatedAt time.Time `gorm:"created_at:time"` //创建时间
}

func UserRegister(user *User) error {
	user.Pwd = tools.PasswordEncode(user.Pwd)
	return DB.Create(user).Error
}

func FindUser(username, passowrd string) *User {
	user := new(User)
	DB.Where("username = ? && pwd = ?", username, passowrd).Find(user)
	return user
}

func FindUserByUserName(username string) *User {
	user := new(User)
	DB.Where("username = ?", username).Find(user)
	return user
}
