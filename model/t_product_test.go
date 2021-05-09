package model

import (
	"fmt"
	"testing"
)

func TestUpdateTProduct(t *testing.T) {
	err := UpdateTProduct(&TProduct{Id: 10001, Name: "测试商品2"})
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
}
