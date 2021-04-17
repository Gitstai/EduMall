package model

import (
	"EduMall/dal"
	"time"
)

type TUser struct {
	Id         int64     `gorm:"column:id" json:"id"`
	Account    string    `gorm:"column:account" json:"account"`
	Password   string    `gorm:"column:password" json:"password"`
	Nickname   string    `gorm:"column:nickname" json:"nickname"`
	UserType   int8      `gorm:"column:user_type" json:"user_type"`
	Balance    int32     `gorm:"column:balance" json:"balance"`
	IsDelete   int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TUser) TableName() string {
	return "t_user"
}

func GetTUserInfo(user *TUser) (res TUser, err error) {
	err = dal.EduDB.Where(user).Where(map[string]interface{}{"is_delete":0}).Find(&res).Error
	return
}


