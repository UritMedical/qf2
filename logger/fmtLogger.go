/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/14 11:23
 */

package logger

import (
	"fmt"
	. "github.com/UritMedical/qf2/define"
	"time"
)

type fmtLogger struct {
	BaseLogger
}

func NewFmt(level ELogLevel) QLogger {

	logger := &fmtLogger{}
	logger.BaseLogger = NewBaseLogger(level, logger.log)
	return logger
}

func (logger *fmtLogger) log(level ELogLevel, format string, params []interface{}) {
	str := fmt.Sprintf("[%v][%s]", time.Now().Format("15:04:05"), level.ToString())
	fmt.Printf(str+format+"\r\n", params...)
}
