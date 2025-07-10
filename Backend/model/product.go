package model

import (
	"time"
)

type Product struct {
	Id             int64     `gorm:"column:ID;type:bigint(20);AUTO_INCREMENT;NOT NULL" json:"ID"`
	ProductName    string    `gorm:"column:product_name;type:varchar(255)" json:"product_name"`
	ProductImage   string    `gorm:"column:product_image;type:varchar(255)" json:"product_image"`
	ProductDetail  string    `gorm:"column:product_detail;type:varchar(255)" json:"product_detail"`
	Price          float64   `gorm:"column:price;type:decimal(10,2)" json:"price"`
	TotalSell      string    `gorm:"column:total_sell;type:varchar(255)" json:"total_sell"`
	Remain         int64     `gorm:"column:remain;type:bigint(20)" json:"remain"`
	Status         int       `gorm:"column:status;type:tinyint(4)" json:"status"`
	ProductAddress string    `gorm:"column:product_address;type:varchar(255)" json:"product_address"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (m *Product) TableName() string {
	return "product"
}
