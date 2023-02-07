package dao

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
database: passby_gwf
username: h2cghemzuz96xd4c50nm
host: ap-northeast.connect.psdb.cloud
password: pscale_pw_Lb3RUI6arn9XmEv2wIfAWdsRJgFya39YrpQjvcazlKv
*/
var DB *gorm.DB

func init() {
	dsn := "h2cghemzuz96xd4c50nm:pscale_pw_Lb3RUI6arn9XmEv2wIfAWdsRJgFya39YrpQjvcazlKv@tcp(ap-northeast.connect.psdb.cloud)/passby_gwf?tls=true&&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SkipDefaultTransaction: true, //关闭事务自动提交
	})
	if err != nil {
		log.Panicln(err)
	} else {
		DB = db
	}
}
