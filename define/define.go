package define

type ApiHandler func(params ...interface{}) (interface{}, error)

type NoticeHandler func(params ...interface{})

type QBll interface {
	//
	// Stop
	//  @Description: 释放资源
	//
	Stop()
	//
	// GetConfig
	//  @Description: 返回业务的conf内容
	//
	GetConfig() string
	//
	// Apis
	//  @Description: 注册API方法
	//  @return map[string]ApiHandler
	//
	Apis() map[string]ApiHandler
	//
	// Subscribes
	//  @Description: 注册订阅方法
	//  @return map[string]NoticeHandler
	//
	Subscribes() map[string]NoticeHandler
	//
	// SendNotice
	//  @Description: 发送消息
	//  @param topic
	//  @param params
	//  @return error
	//
	SendNotice(topic string, params ...interface{}) error
}

type QPlugin interface {
	//
	// GetId
	//  @Description: 返回插件ID
	//  @return string
	//
	GetId() string
	//
	// GetVersion
	//  @Description: 返回插件版本
	//  @return string
	//
	GetVersion() string // 返回插件版本号
	//
	// BuildDocument
	//  @Description: 生成插件conf文档
	//
	BuildDocument()
	//
	// RegBll
	//  @Description: 注册业务
	//  @param bll
	//
	RegBll(bll ...QBll)
	//
	// BindWrapperInvoke
	//  @Description: 绑定打包器的外部访问方法
	//  @param invoke 方法指针
	//
	BindWrapperInvoke(invoke func(plugin string, route string, params ...interface{}) (interface{}, error))
	//
	// BindWrapperNotice
	//  @Description: 绑定打包器的发送的消息方法
	//  @param invoke 方法指针
	//
	BindWrapperNotice(notice func(topic string, params ...interface{}) error)
	//
	// Invoke
	//  @Description: 执行方法
	//  @param route
	//  @param params
	//  @return interface{}
	//  @return error
	//
	Invoke(route string, params ...interface{}) (interface{}, error)
}
