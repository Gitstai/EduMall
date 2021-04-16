package model

import "time"

type TProductDescription struct {
	Id         int64     `gorm:"column:id" json:"id"`
	DescText   string    `gorm:"column:desc_text" json:"desc_text"`
	DescImg    string    `gorm:"column:desc_img" json:"desc_img"`
	IsDelete   int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TProductDescription) TableName() string {
	return "t_product_description"
}
