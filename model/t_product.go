package model

import "time"

type TProduct struct {
	Id            int64     `gorm:"column:id" json:"id"`
	ProviderId    int64     `gorm:"column:provider_id" json:"provider_id"`
	Name          string    `gorm:"column:name" json:"name"`
	ProductType   int8      `gorm:"column:product_type" json:"product_type"`
	Status        int8      `gorm:"column:status" json:"status"`
	Price         int32     `gorm:"column:price" json:"price"`
	Keywords      string    `gorm:"column:keywords" json:"keywords"`
	DescId        int64     `gorm:"column:desc_id" json:"desc_id"`
	FileId        int64     `gorm:"column:file_id" json:"file_id"`
	AfterSaleText string    `gorm:"column:after_sale_text" json:"after_sale_text"`
	Inventory     int32     `gorm:"column:inventory" json:"inventory"`
	SaleVolume    int32     `gorm:"column:sale_volume" json:"sale_volume"`
	IsDelete      int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TProduct) TableName() string {
	return "t_product"
}
