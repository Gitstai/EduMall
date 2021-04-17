package logs

import (
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"strconv"
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger()
	i := 1
	logger.Debug("Debug>>>>>>>>>>>>>>>>>>>>>>", strconv.Itoa(i))
	logger.Info("Info>>>>>>>>>>>>>>>>>>>>>>>>>", strconv.Itoa(i))
	logger.Warn(fmt.Sprintf("Warn>>>>>>>>>>>>>>>>>>>>>>>>>%v", strconv.Itoa(i)))
	logger.Error("Error>>>>>>>>>>>>>>>>>>>>>>>>>", strconv.Itoa(i))
	logger.Fatal("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>", strconv.Itoa(i))
	logger.Info("resp no ok st:", 12, "msg:", 32)
}

func TestName(t *testing.T) {
	try := make(map[string]interface{})
	try["sub"] = int64(1)
	switch try["sub"].(type) {
	case float64:
		fmt.Println("float64")
	case int64:
		fmt.Println("int64")
	default:
		fmt.Println("other")
	}
	a := try["sub"].(int64)
	fmt.Println(a)
}
