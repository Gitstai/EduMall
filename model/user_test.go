package model

import (
	"EduMall/dal"
	"EduMall/tools"
	"fmt"
	"testing"
)

func init() {
	err := dal.InitDB()
	if err != nil {
		panic(err)
	}
}

func TestGetUserInfo(t *testing.T) {
	info, err := GetUserInfo(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(tools.ToJson(info))
}