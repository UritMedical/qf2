/**
 * @Author: Joey
 * @Description:接口声明
 * @Create Date: 2023/7/26 12:07
 */

package define

//
// QBus
// @Description: 总线接口
//
type QBus interface {
	//Plug 接入插件
	Plug(QPlugin) []QError
	Invoke(route string, args map[string]interface{}) (interface{}, QError)
	SendNotice(topic string, isSync bool, arg interface{}) QError
	Logger() QLogger
}

//
// QPlugin
// @Description: 插件接口
//
type QPlugin interface {

	//RegBus 注册总线
	RegBus(QBus)
	//Apis 向总线提供API声明
	Apis() map[string]ApiHandler
	//Subscribes 向总线提供消息订阅声明
	Subscribes() map[string]NoticeHandler
	Init() QError
	Stop() QError
	//GetId 向总线提供业务唯一编号 自动化测试和log需求
	GetId() string
	//Invoke 调用总线上的业务方法
	Invoke(route string, args map[string]interface{}) (interface{}, QError)
	//SendNotice 向总线上的业务方法
	SendNotice(topic string, isSync bool, arg interface{}) QError
	Logger() QLogger
}

type QLogger interface {
	Debug(format string, params ...interface{})
	Info(format string, params ...interface{})
	Warn(format string, params ...interface{})
	Error(format string, params ...interface{})
	Fatal(format string, params ...interface{})
}
