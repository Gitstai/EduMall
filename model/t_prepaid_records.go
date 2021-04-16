package model

import "time"

type TPrepaidRecords struct {
	Id         int64     `gorm:"column:id" json:"id"`
	UserId     int64     `gorm:"column:user_id" json:"user_id"`
	CdKey      string    `gorm:"column:cd_key" json:"cd_key"`
	Amount     int32     `gorm:"column:amount" json:"amount"`
	IsDelete   int8     `gorm:"column:is_delete" json:"is_delete"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TPrepaidRecords) TableName() string {
	return "t_prepaid_records"
}
