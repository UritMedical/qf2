/**
 * @Author: Joey
 * @Description:枚举 函数指针等的定义
 * @Create Date: 2023/7/26 13:58
 */

package define

import (
	"gorm.io/gorm"
)

// 注册相关函数定义
type (
	ApiHandler    func(map[string]interface{}) (interface{}, QError) // API函数指针
	NoticeHandler func(interface{})                                  // 消息函数指针
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

//
// QError
//  @Description: 异常
//
type QError interface {
	//
	// Code
	//  @Description: 获取故障码
	//  @return int
	//
	Code() int
	//
	// Error
	//  @Description: 获取异常描述
	//  @return string
	//
	Error() string
}

const (
	ErrorCodeParamInvalid        = iota + 100 // 传入参数无效
	ErrorCodePermissionDenied                 // 权限不足，拒绝访问
	ErrorCodeRecordNotFound                   // 未找到记录
	ErrorCodeRecordExist                      // 记录已经存在
	ErrorCodeSaveFailure                      // 保存失败
	ErrorCodeDeleteFailure                    // 删除失败
	ErrorCodeFileNotFound                     // 文件不存在
	ErrorCodeUploadedFileNull                 // 未上传任何文件
	ErrorCodeUploadedFileInvalid              // 上传文件解析失败
	ErrorCodeTImeOut                          // 超时
	ErrorCodeAPIUndefined                     // API未定义
	ErrorCodeNoticeUndefined                  // 消息未定义
)

const (
	ErrorCodeOSError = 900 // 系统故障
	ErrorCodeUnknown = 999 // 未知异常
)

var errorCodeTextMap = map[int]string{
	ErrorCodeParamInvalid:        "无效的参数",
	ErrorCodePermissionDenied:    "权限不足，拒绝访问",
	ErrorCodeRecordNotFound:      "未找到记录",
	ErrorCodeRecordExist:         "记录已经存在",
	ErrorCodeSaveFailure:         "保存失败",
	ErrorCodeDeleteFailure:       "删除失败",
	ErrorCodeFileNotFound:        "指定文件不存在",
	ErrorCodeUploadedFileNull:    "未上传任何文件",
	ErrorCodeUploadedFileInvalid: "上传文件解析失败",
	ErrorCodeOSError:             "系统运行故障",
	ErrorCodeUnknown:             "其他未知故障",
	ErrorCodeTImeOut:             "请求超时",
	ErrorCodeAPIUndefined:        "API未定义",
	ErrorCodeNoticeUndefined:     "消息未定义",
}

//
// Error
//  @Description: 创建故障内容
//  @param code
//  @param err
//  @return IError
//
func Error(code int, err string) QError {
	return errorInfo{
		code:  code,
		error: err,
	}
}

type errorInfo struct {
	code  int
	error string
}

func (e errorInfo) Code() int {
	return e.code
}

func (e errorInfo) Error() string {
	return e.error
}

type QDb struct {
	DB *gorm.DB
}
