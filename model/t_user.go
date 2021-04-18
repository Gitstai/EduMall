package model

import (
	"EduMall/dal"
	"errors"
	"time"
)

type TUser struct {
	Id         int64     `gorm:"column:id" json:"id"`
	Account    string    `gorm:"column:account" json:"account"`
	Password   string    `gorm:"column:password" json:"password"`
	Nickname   string    `gorm:"column:nickname" json:"nickname"`
	UserType   int8      `gorm:"column:user_type;default:2" json:"user_type"`
	Balance    int32     `gorm:"column:balance" json:"balance"`
	IsDelete   int8      `gorm:"column:is_delete" json:"is_delete"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TUser) TableName() string {
	return "t_user"
}

func GetTUserInfo(user *TUser) ([]*TUser, error) {
	var res []*TUser
	err := dal.EduDB.Where(user).Where(map[string]interface{}{"is_delete": 0}).Find(&res).Error
	return res, err
}

func InsertTUser(data *TUser) (*TUser, error) {
	if data == nil {
		return data, errors.New("insert no data")
	}
	err := dal.EduDB.Create(data).Error
	return data, err
}
