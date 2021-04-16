package model

import (
	"fmt"
	"testing"
)

func TestGetTUserInfo(t *testing.T) {
	info, err := GetTUserInfo(&TUser{Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}