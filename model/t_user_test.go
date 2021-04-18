package model

import (
	"EduMall/tools"
	"fmt"
	"testing"
)

func TestGetTUserInfo(t *testing.T) {
	info, err := GetTUserInfo(&TUser{Id: 10001})
	if err != nil {
		panic(err)
	}
	fmt.Println(len(info))
}

func TestInsertTUser(t *testing.T) {
	u, err := InsertTUser(&TUser{Account: "15200830961", Nickname: "MemoKun2", Password: "123"})
	if err != nil {
		panic(err)
	}
	fmt.Println(tools.ToJson(u))
}
