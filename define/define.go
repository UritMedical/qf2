/**
 * @Author: Joey
 * @Description:枚举 函数指针等的定义
 * @Create Date: 2023/7/26 13:58
 */

package define

// 注册相关函数定义
type (
	ApiHandler    func(map[string]interface{}) (interface{}, error) // API函数指针
	NoticeHandler func(interface{})                                 // 消息函数指针
)

type ELogLevel int

const (
	ELogLevelDebug = 1
	ELogLevelInfo  = 2
	ELogLevelWarn  = 4
	ELogLevelError = 8
	ELogLevelFatal = 16
)

func (level ELogLevel) ToString() string {
	switch level {
	case ELogLevelDebug:
		return "debug"
	case ELogLevelInfo:
		return "info"
	case ELogLevelWarn:
		return "warn"
	case ELogLevelError:
		return "error"
	case ELogLevelFatal:
		return "fatal"
	default:
		return "unknown"
	}
}
