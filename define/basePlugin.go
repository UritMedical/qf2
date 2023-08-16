/**
 * @Author: Joey
 * @Description:基类 简化继承自接口需要写的代码
 * @Create Date: 2023/8/14 16:10
 */

package define

type QBasePlugin struct {
	bus QBus
}

func (plugin *QBasePlugin) RegBus(bus QBus) {
	plugin.bus = bus
}
func (plugin *QBasePlugin) Invoke(route string, args map[string]interface{}) (interface{}, QError) {
	return plugin.bus.Invoke(route, args)
}

func (plugin *QBasePlugin) SendNotice(topic string, isSync bool, arg interface{}) QError {
	return plugin.bus.SendNotice(topic, isSync, arg)
}
func (plugin *QBasePlugin) Init() QError {
	return nil
}
func (plugin *QBasePlugin) Stop() QError {
	return nil
}
func (plugin *QBasePlugin) Logger() QLogger {
	return plugin.bus.Logger()
}

//
// BaseModel
//  @Description: 基础实体对象
//
type BaseModel struct {
	Id       uint64   `gorm:"primaryKey"` // 唯一号
	LastTime DateTime `gorm:"index"`      // 最后操作时间时间
	Summary  string   // 摘要
	FullInfo string   // 其他扩展内容
}
