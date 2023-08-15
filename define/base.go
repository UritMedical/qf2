/**
 * @Author: Joey
 * @Description:基类 简化继承自接口需要写的代码
 * @Create Date: 2023/8/14 16:10
 */

package define

type QBasePlugin struct {
	QPlugin
	bus QBus
}

func (plugin *QBasePlugin) RegBus(bus QBus) {
	plugin.bus = bus
}
func (plugin *QBasePlugin) Invoke(route string, args map[string]interface{}) (interface{}, error) {
	return plugin.bus.Invoke(route, args)
}

func (plugin *QBasePlugin) SendNotice(topic string, isSync bool, arg interface{}) error {
	return plugin.bus.SendNotice(topic, isSync, arg)
}
func (plugin *QBasePlugin) Init() error {
	return nil
}
func (plugin *QBasePlugin) Stop() error {
	return nil
}
func (plugin *QBasePlugin) Logger() QLogger {
	return plugin.bus.Logger()
}
