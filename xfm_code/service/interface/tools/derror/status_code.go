package derror

import (
	"tools/gerror"
)

const (
	SuccessCode  = 0
	DbErrCode    = 601
	NilErrCode   = 602
	ParaErrCode  = 603
	LogicErrCode = 604
	OtherErrCode = 605
	PanicErrCode = 606
)

// ResponseDBErr 操作db相关错误
func ResponseDBErr(data interface{}, msg string) (interface{}, error) {
	return data, BuildErrMsg(DbErrCode, msg)
}

// ResponseNilErr 空错误
func ResponseNilErr(data interface{}, msg string) (interface{}, error) {
	return data, BuildErrMsg(NilErrCode, msg)
}

// ResponseParaErr 参数校验错误
func ResponseParaErr(data interface{}, msg string) (interface{}, error) {
	return data, BuildErrMsg(ParaErrCode, msg)
}

// ResponseLogicErr 逻辑处理错误
func ResponseLogicErr(data interface{}, msg string) (interface{}, error) {
	return data, BuildErrMsg(LogicErrCode, msg)
}

// ResponseOtherErr 其他错误
func ResponseOtherErr(data interface{}, msg string) (interface{}, error) {
	return data, BuildErrMsg(OtherErrCode, msg)
}

// ResponsePanicErr panic错误
func ResponsePanicErr(data interface{}, msg string) (interface{}, error) {
	return data, BuildErrMsg(PanicErrCode, msg)
}

// ResponseMsgErr 自定义错误码返回
func ResponseMsgErr(code int32, data interface{}, msg string) (interface{}, error) {
	return data, BuildErrMsg(code, msg)
}

// ResponseErr 自定义错误返回
func ResponseErr(data interface{}, err error) (interface{}, error) {
	return data, err
}

// ResponseSuccess 正常返回
func ResponseSuccess(data interface{}) (interface{}, error) {
	return data, BuildErrMsg(SuccessCode, "success")
}

// BuildErrMsg 构建错误
func BuildErrMsg(code int32, msg string) *gerror.Error {
	return &gerror.Error{Code: code, Msg: msg}
}
