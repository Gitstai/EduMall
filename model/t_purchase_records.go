package model

import "time"

type TPurchaseRecords struct {
	Id         int64     `gorm:"column:id" json:"id"`
	UserId     int64     `gorm:"column:user_id" json:"user_id"`
	ProductId  int64     `gorm:"column:product_id" json:"product_id"`
	Payment    int32     `gorm:"column:payment" json:"payment"`
	IsDelete   int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TPurchaseRecords) TableName() string {
	return "t_purchase_records"
}
