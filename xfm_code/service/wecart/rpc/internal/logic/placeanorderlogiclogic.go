package logic

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"runtime/debug"
	"tools/derror"

	"wecart/rpc/internal/svc"
	"wecart/rpc/wecart"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceAnOrderLogicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceAnOrderLogicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceAnOrderLogicLogic {
	return &PlaceAnOrderLogicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlaceAnOrderLogicLogic) PlaceAnOrderLogic(in *wecart.PlaceAnOrderReq) (result *wecart.PlaceAnOrderResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			logx.Error(err, "OrderDetailsLogic_recover()", "in", in)
			debug.PrintStack()
		}
	}()

	ctx, client := SendWeCart()

	svc := jsapi.JsapiApiService{Client: client}
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, APIResult, err := svc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String("wxd678efh567hg6787"),
			Mchid:       core.String("1900009191"),
			Description: core.String("Image形象店-深圳腾大-QQ公仔"),
			OutTradeNo:  core.String("1217752501201407033233368018"),
			Attach:      core.String("自定义数据说明"),
			NotifyUrl:   core.String("https://www.weixin.qq.com/wxpay/pay.php"),
			Amount: &jsapi.Amount{
				Total: core.Int64(100),
			},
			Payer: &jsapi.Payer{
				Openid: core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
			},
		},
	)

	if err != nil {
		obj, err := derror.ResponseMsgErr(int32(APIResult.Response.StatusCode), resp, "query wecart failed")
		result = InterfaceToStruct(obj)
		return result, err
	}
	result.Data[0].PrepayId = *resp.PrepayId
	result.Data[0].Appid = *resp.Appid
	result.Data[0].TimeStamp = *resp.TimeStamp
	result.Data[0].NonceStr = *resp.NonceStr
	result.Data[0].Package = *resp.Package
	result.Data[0].SignType = *resp.SignType
	result.Data[0].PaySign = *resp.PaySign

	obj, err := derror.ResponseSuccess(result)
	result = InterfaceToStruct(obj)
	return result, err
}

// InterfaceToStruct interface转任意类型，每个logic.go中返回类型`*modulation.OrderDetailsResp`记得替换成logic函数的返回类型
func InterfaceToStruct(object interface{}) *wecart.PlaceAnOrderResp {
	obj := object.(*wecart.PlaceAnOrderResp)
	return obj
}
