# qf2 Guid

总线和内存插件的基础使用

请参考test\directBusTest文件

## Step1 get项目

```go
go get github.com/UritMedical/qf2

```

如果需要拉指定版本

 ```go
 go get github.com/UritMedical/qf2/@v0.0.1
 ```

## Step2 定义插件

可以直接实现QPlugin，但是要写的东西就比较多

推荐的方法是包裹QBasePlugin

```
import (
	qf "github.com/UritMedical/qf2"
	. "github.com/UritMedical/qf2/define"
)

type <插件名> struct {
   QBasePlugin
}

func (p *Plugin) Apis() map[string]ApiHandler {
	return nil
}

func (p *Plugin) Subscribes() map[string]NoticeHandler {
	return nil
}

func (p *Plugin) GetId() string {
	return "TestPlugin"
}
```
#Step3 创建总线并挂载插件

```go
bus := qf2.Bus.NewDirect(qf2.Logger.NewFmt()) // 运行插件
bus.Plug(&Plugin{})                            // 注册插件
```