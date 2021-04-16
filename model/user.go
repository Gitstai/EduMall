package model

import (
	"EduMall/dal"
)

type User struct {
	Id   int64  `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName sets the insert table name for this struct type
func (t *User) TableName() string {
	return "user"
}

func GetUserInfo(u *User) (res User, err error) {
	err = dal.EduDB.Where(u).Find(&res).Error
	return res, err
}
