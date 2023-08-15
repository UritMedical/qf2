/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/14 16:50
 */

package logger

import (
	"encoding/json"
	"fmt"
	. "github.com/UritMedical/qf2/define"
	"github.com/UritMedical/qf2/util/mqtt"
	"time"
)

type mqttLogger struct {
	BaseLogger
	client *mqtt.Client
	topic  string
}

func NewMQTT(level ELogLevel, topic, host string) QLogger {

	logger := &mqttLogger{}
	logger.topic = topic
	logger.BaseLogger = NewBaseLogger(level, logger.log)
	logger.client = mqtt.NewClient(host, time.Now().String())
	return logger
}

func (logger *mqttLogger) log(level ELogLevel, format string, params []interface{}) {
	info := fmt.Sprintf(format, params...)
	log := struct {
		Time  time.Time
		Level string
		Info  string
	}{
		time.Now(),
		level.ToString(),
		info,
	}
	bytes, _ := json.Marshal(log)

	//str := fmt.Sprintf("[%v][%s]", time.Now().Format("15:04:05"), level.ToString())
	//str = fmt.Sprintf(str+format+"\r\n", params...)
	logger.client.Publish(logger.topic, 0, bytes, false)
}
