package qlog

import (
	"fmt"
	"github.com/UritMedical/qf2/utils/qdate"
	"github.com/UritMedical/qf2/utils/qio"
	"time"
)

// Debug
//
//	@Description: 调试日志
//	@param content 内容
func Debug(content string) {
	writeLog(content, "Debug")
}

// Warn
//
//	@Description: 警告日志
//	@param content 内容
func Warn(content string) {
	writeLog(content, "Warn")
}

// Error
//
//	@Description: 异常日志
//	@param content 内容
//	@param err 异常
func Error(content string, err error) {
	if err != nil {
		content += err.Error()
	}
	writeLog(content, "Error")
}

func writeLog(msg string, tp string) {
	logStr := fmt.Sprintf("DateTime: %s\n", qdate.ToString(time.Now(), "yyyy-MM-dd HH:mm:ss"))
	logStr += fmt.Sprintf("Messages: %s\n", msg)
	logStr += "----------------------------------------------------------------------------------------------\n\n"
	per := qdate.ToString(time.Now(), "yyyy-MM")
	day := qdate.ToString(time.Now(), "dd")
	logFile := fmt.Sprintf("./log/%s/%s_%s.log", per, day, tp)
	logFile = qio.GetFullPath(logFile)
	_ = qio.WriteString(logFile, logStr, true)
}
