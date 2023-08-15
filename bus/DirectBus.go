/**
 * @Author: Joey
 * @Description:直通型总线
 * @Create Date: 2023/7/26 17:49
 */

package bus

import (
	"errors"
	. "github.com/UritMedical/qf2/define"
)

//
// directBus
// @Description: 直通型总线 内存级直接通信
//
type directBus struct {
	apiDict    map[string]ApiHandler
	noticeDict map[string][]NoticeHandler
	logger     QLogger
}

func (bus *directBus) Logger() QLogger {
	return bus.logger
}
func (bus *directBus) Plug(plugin QPlugin) []error {
	var errs []error
	// 注册API
	apis := plugin.Apis()
	for k, v := range apis {
		if err := bus.regApi(k, v); err != nil {
			errs = append(errs, err...)
		}
	}
	// 注册消息订阅
	subscribes := plugin.Subscribes()
	for k, v := range subscribes {
		bus.subscribe(k, v)
	}
	return errs
}

func NewDirect(logger QLogger) QBus {
	bus := &directBus{
		apiDict:    make(map[string]ApiHandler),
		noticeDict: make(map[string][]NoticeHandler),
		logger:     logger,
	}
	return bus
}

func (bus *directBus) regApi(k string, v ApiHandler) []error {
	var errs []error
	_, b := bus.apiDict[k]
	if b {
		errs = append(errs, errors.New("api already exists")) // 重复注册
	}
	bus.apiDict[k] = v // 注册API
	return errs
}

func (bus *directBus) subscribe(k string, v NoticeHandler) {
	_, b := bus.noticeDict[k]
	if !b {
		bus.noticeDict[k] = make([]NoticeHandler, 0)
	}
	bus.noticeDict[k] = append(bus.noticeDict[k], v) // 注册消息
}

func (bus *directBus) Invoke(route string, args map[string]interface{}) (interface{}, error) {
	function, b := bus.apiDict[route]
	if !b {
		bus.logger.Error("invoke api failed api %s not found", route)
		return nil, errors.New("api not found")
	}
	return function(args) // 调用API
}

func (bus *directBus) SendNotice(topic string, isSync bool, arg interface{}) error {
	notice, b := bus.noticeDict[topic]
	if !b {
		bus.logger.Error("send notice failed notice %s not found", topic)
		return errors.New("notice not found")
	}
	if isSync {
		for _, handler := range notice {
			handler(arg)
		}
	} else {
		for _, handler := range notice {
			go handler(arg)
		}
	}
	return nil // 发送消息
}
