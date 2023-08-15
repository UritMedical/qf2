/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/12 10:20
 */

package qf2

import (
	"github.com/UritMedical/qf2/bus"
	. "github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/logger"
)

func Run(bus QBus, plugins []QPlugin) []error {

	var errs []error
	for _, v := range plugins {
		errs = append(errs, v.Init())
		v.RegBus(bus)
		errs = append(errs, bus.Plug(v)...)
	}
	return errs
}

type busFactory struct {
}

var Bus *busFactory

func (_ *busFactory) NewDirect(logger QLogger) QBus {
	b := bus.NewDirect(logger) // 运行插件
	return b
}

type loggerFactory struct {
}

var Logger *loggerFactory

func (_ *loggerFactory) NewFmt() QLogger {
	return logger.NewFmt(ELogLevelDebug | ELogLevelInfo | ELogLevelWarn | ELogLevelError)
}
func (_ *loggerFactory) NewMQTT() QLogger {
	return logger.NewMQTT(ELogLevelDebug|ELogLevelInfo|ELogLevelWarn|ELogLevelError, "Q-LOG", "127.0.0.1:1883")
}
