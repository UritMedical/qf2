/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/8/12 13:46
 */

package main

import (
	"fmt"
	qf "github.com/UritMedical/qf2"
	. "github.com/UritMedical/qf2/define"
)

func main() {
	plugins := []QPlugin{
		NewPluginA("PluginA"),
		NewPluginB("PluginB"),
	}
	//bus := qf.Bus.NewDirect(qf.Logger.NewFmt())
	bus := qf.Bus.NewDirect(qf.Logger.NewMQTT())
	qf.Run(bus, plugins) // 运行插件
	fmt.Println(plugins[0].Invoke("Hello", map[string]interface{}{"input": "input From Main"}))
	fmt.Println(plugins[0].Invoke("Hello1", map[string]interface{}{"input": "input From Main"}))
	fmt.Println(plugins[1].SendNotice("topic", true, "topic From Main"))   // 发送消息
	fmt.Println(plugins[1].SendNotice("topic1", true, "topic1 From Main")) // 发送消息
	fmt.Println(plugins[0].Invoke("GoodBye", map[string]interface{}{"input": "input From Main"}))
	fmt.Println(plugins[1].Invoke("Hello", map[string]interface{}{"input": "input From Main"}))
	fmt.Println(plugins[1].Invoke("GoodBye", map[string]interface{}{"input": "input From Main"}))
	bus.Logger().Error("123")
	a := "a"
	_, err := fmt.Scan(&a)
	if err != nil {
		return
	}
}
