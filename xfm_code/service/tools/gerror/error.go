package gerror

import "fmt"

// Error 错误结构体
type Error struct {
	Code int32
	Msg  string
	Detail string
}

// Error 返回错误信息
func (e *Error) Error() string {
	return e.Msg
}

// ErrorInfo 详细错误信息
func (e *Error) ErrorInfo() (int32, error) {
	return e.Code, fmt.Errorf(e.Msg)
}

// IsError 是否出错
func (e *Error) IsError() bool {
	return e.Code != 0 || e.Msg != ""
}

// NewError 创建错误对象
func NewError(code int32, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

// NewErrorDetail 创建错误对象
func NewErrorDetail(err *Error, detailMsg string, a...interface{}) *Error {
	errDesc := fmt.Sprintf(detailMsg, a...)
	return &Error{
		Code: err.Code,
		Msg: err.Msg,
		Detail: errDesc,
	}
}

// GetRawError 获取原始错误对象
func GetRawError(err error) *Error {
	ee, ok := err.(*Error)
	if ok {
		return ee
	}
	return &Error{
		Code:9999,
		Msg: fmt.Sprintf("UnKnow error : %v", err)}
}

// GetErrCode 获取errCode
func GetErrCode(err error) int32 {
	return GetRawError(err).Code
}
