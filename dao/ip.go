package dao

import "time"

type Ip struct {
	ID        uint      `gorm:"primary_key"`
	IpName    string    //文件名称
	DDL       string    //截止日期
	CreatedAt time.Time `gorm:"created_at:time"`
}

func CreateIp(fileName, ddl string) error {
	ip := &Ip{IpName: fileName, DDL: ddl}
	result := DB.Create(ip)
	return result.Error
}

func FindIpByDDL(ddl string) *Ip {
	ip := new(Ip)
	DB.Where("ddl = ?", ddl).Find(ip)
	return ip
}

func FindIpFirst() *Ip {
	i := new(Ip)
	DB.First(i)
	return i
}
