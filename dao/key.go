package dao

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Key struct {
	ID            uint `gorm:"primary_key"`
	Key           string
	Traffic       int       `gorm:"default:10240"` //流量 单位M
	ValidDuration int       `gorm:"default:30"`    // 有效时长 单位：天
	Used          bool      `gorm:"default:false"`
	CreatedAt     time.Time `gorm:"created_at:time"`
}

type UserTraffic struct {
	ID        uint   `gorm:"primary_key"`
	UserName  string `gorm:"column:username"` //用户名
	Key       string
	Traffic   int //流量 单位M
	Deadline  time.Time
	CreatedAt time.Time `gorm:"created_at:time"`
}

func (UserTraffic) TableName() string {
	return "users_traffic"
}

/*
	获取未使用密钥
*/
func FindKeys() []Key {
	keys := make([]Key, 0)
	DB.Find(&keys)
	return keys
}

/*
	随机创建密钥，并存入数据库中
	count 创建密钥的个数
*/
func CreateKey(count int) error {
	keys := make([]Key, count)

	for i := 0; i < count; i++ {
		u, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		keys[i] = Key{Key: u.String()}
	}

	result := DB.Create(keys)
	return result.Error
}

/*
	密钥绑定设备
	1.密钥绑定后未到期不可解绑
	2.账号流量扣除优先扣除最先过期那条数据的流量
*/
func BindKey(keyStr, username string) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		key := new(Key)
		tx.Where("`key`=? && used=false", keyStr).Find(key)
		if key.ID == 0 {
			return fmt.Errorf("无效兑换码:%s", keyStr)
		}
		user := new(User)
		tx.Where("username = ?", username).Find(user)
		if user.ID == 0 {
			return fmt.Errorf("用户不存在：%s", username)
		}
		if err := tx.Model(key).Update("used", true).Error; err != nil {
			return err
		}

		userTraffic := &UserTraffic{UserName: username, Key: keyStr, Traffic: key.Traffic, Deadline: time.Now().Add(time.Hour * 24 * time.Duration(key.ValidDuration))}
		if err := tx.Create(userTraffic).Error; err != nil {
			return err
		}
		return nil
	})
}

/*
	流量扣除
*/
func TrafficDeduction(ut UserTraffic) error {

	return nil
}

/*
	查询单个用户流量
*/
func FindTrafficByUser(username string) ([]UserTraffic, error) {
	ut := make([]UserTraffic, 0)
	return ut, DB.Where("username = ?", username).Order("deadline").Find(&ut).Error
}

/*
	查询全部用户流量
*/
