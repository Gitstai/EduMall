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
