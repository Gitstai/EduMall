package model

import "time"

type TProductFile struct {
	Id         int64     `gorm:"column:id" json:"id"`
	FileType   int8      `gorm:"column:file_type" json:"file_type"`
	FileName   string    `gorm:"column:file_name" json:"file_name"`
	Url        string    `gorm:"column:url" json:"url"`
	IsDelete   int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TProductFile) TableName() string {
	return "t_product_file"
}
