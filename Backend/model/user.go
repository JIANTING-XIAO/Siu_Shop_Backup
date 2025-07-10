package model

import "time"

type User struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Account     string    `gorm:"column:account;default:NULL"`
	Password    string    `gorm:"column:password;default:NULL"`
	Name        string    `gorm:"column:name;default:NULL"`
	Telephone   string    `gorm:"column:telephone;default:NULL"`
	Email       string    `gorm:"column:email;default:NULL"`
	Avatar      string    `gorm:"column:avatar;type:mediumblob" `
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UpdateTime  time.Time `gorm:"column:update_time;default:NULL"`
}

// TableName 表名
func (u *User) TableName() string {

	return "user"
}
