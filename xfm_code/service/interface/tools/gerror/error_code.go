package gerror

var (
	ErrCKV                      = Error{Code: 1, Msg: "CKV错误"}
	ErrUpdateUnMatch            = Error{Code: 2, Msg: "Mongo更新CAS冲突"}
	ErrReqParamInvalid          = Error{Code: 3, Msg: "请求参数格式错误"}
	ErrPulsarConf               = Error{Code: 4, Msg: "Pulsar配置错误"}
	ErrPulsarClientCreateFailed = Error{Code: 5, Msg: "Pulsar客户端创建失败"}
	ErrDataParse                = Error{Code: 5, Msg: "数据解析失败"}
	ErrPulsarProduceCreate      = Error{Code: 6, Msg: "pulsarProducer创建失败"}
	ErrPulsarSend               = Error{Code: 7, Msg: "pulsarProducer发送失败"}
	ErrDbQuery                  = Error{Code: 8, Msg: "数据库查询错误"}
	ErrEs                       = Error{Code: 9, Msg: "数据库操作错误"}
	ErrTdwApi                   = Error{Code: 10, Msg: "天穹元数据接口调用返回异常"}
	ErrJson                     = Error{Code: 11, Msg: "json系列化出错"}
	ErrTime                     = Error{Code: 12, Msg: "非法的时间格式"}
	ErrRedis                    = Error{Code: 13, Msg: "redis请求出错"}
	ErrIoaRPC                   = Error{Code: 14, Msg: "ioa验证远程调用出错"}
	ErrIoaFormatError           = Error{Code: 15, Msg: "ioa_token失效或错误, 请重新登录"}
)
