package utils

import (
	"fmt"
	"runtime"
	"time"
)

func writeLogFile(level string, log string, callerLevel int) {
	// write now
	_, filename, line, _ := runtime.Caller(callerLevel)
	_, err := SelfLogWriter.WriteString(fmt.Sprintf("%s, %s\tfile:///%s:%d\t %s%s\n", level,
		time.Now().Format("2006-01-02 15:04:05"),
		filename, line,
		log, Resets))
	if err != nil {
		println(err.Error())
	}
}

const (
	green  = "\033[32m" // 绿色
	red    = "\033[31m" // 红色
	yellow = "\033[33m" // 黄色
	Resets = "\033[0m"  // 重置颜色
)

func LogInfo(log string) {
	writeLogFile(fmt.Sprintf("%s%s", green, "INFO"), log, 2)
}

func LogWarn(log string) {
	writeLogFile(fmt.Sprintf("%s%s", yellow, "WARN"), log, 2)
}

func LogError(log string) {
	writeLogFile(fmt.Sprintf("%s%s", red, "ERROR"), log, 2)
}
