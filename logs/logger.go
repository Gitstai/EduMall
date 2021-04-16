package logs

import (
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"time"
)

func InitLogger() {
	//制定是否控制台打印 默认为true
	logger.SetConsole(true)

	//指定日志文件备份方式为日期的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	fileName := fmt.Sprintf(time.Now().Format("2006_01_02")+".log")
	logger.SetRollingDaily("../logs", fileName)

	//指定日志级别 ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	logger.SetLevel(logger.DEBUG)
}

func Debug(format string, a ...interface{}) {
	logger.Debug(fmt.Sprintf(format, a...))
}
func Info(format string, a ...interface{}) {
	logger.Info(fmt.Sprintf(format, a...))
}
func Warn(format string, a ...interface{}) {
	logger.Warn(fmt.Sprintf(format, a...))
}
func Error(format string, a ...interface{}) {
	logger.Error(fmt.Sprintf(format, a...))
}
func Fatal(format string, a ...interface{}) {
	logger.Fatal(fmt.Sprintf(format, a...))
}
