package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func writeLogFile(level string, log string, callerLevel int) {
	// write now
	_, filename, line, _ := runtime.Caller(callerLevel)
	_, err := writer.WriteString(fmt.Sprintf("%s, %s\tfile:///%s:%d\t %s%s\n", level,
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

func LogErrorThrough(err interface{}) error {
	var actualErr error

	switch e := err.(type) {
	case error:
		// If it's an error, use it directly.
		actualErr = e
	case *error:
		// If it's a pointer to an error, dereference it.
		if e != nil {
			actualErr = *e
		}
	default:
		// If it's neither, perhaps return nil or handle differently.
		return nil
	}

	if actualErr != nil {
		LogErrorWithCallerLevel(actualErr.Error(), 3)
	}

	return actualErr
}

func LogErrorWithCallerLevel(log string, callerLevel int) {
	writeLogFile("ERROR", log, callerLevel)
}

var logFile *os.File

// default std out
var writer = os.Stdout

func InitLog() (err error) {
	// use LogFileName get path
	if useLogFile != "0" {
		fullFileName := logFilePath
		// 判断是否是绝对路径
		if !filepath.IsAbs(logFilePath) {
			fullFileName = filepath.Join(os.TempDir(), logFilePath)
		}
		dirName := filepath.Dir(fullFileName)
		err = os.MkdirAll(dirName, 0o755)
		if err != nil {
			return err
		}
		logFile, err = os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)

		// 适当的条件下设置writer = logFile,默认是os.Stdout
		writer = logFile
	}

	return err
}
