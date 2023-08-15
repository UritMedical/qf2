/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/14 16:42
 */

package logger

import (
	. "qf/define"
)

type BaseLogger struct {
	level ELogLevel
	log   func(level ELogLevel, format string, params []interface{})
}

func NewBaseLogger(level ELogLevel, l func(level ELogLevel, format string, params []interface{})) BaseLogger {
	return BaseLogger{
		level: level,
		log:   l,
	}
}
func (logger *BaseLogger) Debug(format string, params ...interface{}) {
	if logger.level&ELogLevelDebug == ELogLevelDebug {
		logger.log(ELogLevelDebug, format, params)
	}
}

func (logger *BaseLogger) Info(format string, params ...interface{}) {
	if logger.level&ELogLevelInfo == ELogLevelInfo {
		logger.log(ELogLevelInfo, format, params)
	}
}

func (logger *BaseLogger) Warn(format string, params ...interface{}) {
	if logger.level&ELogLevelWarn == ELogLevelWarn {
		logger.log(ELogLevelWarn, format, params)
	}
}

func (logger *BaseLogger) Error(format string, params ...interface{}) {
	if logger.level&ELogLevelError == ELogLevelError {
		logger.log(ELogLevelError, format, params)
	}
}

func (logger *BaseLogger) Fatal(format string, params ...interface{}) {
	if logger.level&ELogLevelFatal == ELogLevelFatal {
		logger.log(ELogLevelFatal, format, params)
	}
}
