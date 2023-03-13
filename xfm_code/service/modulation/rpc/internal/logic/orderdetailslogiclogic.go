package logic

import (
	"context"
	"modulation/rpc/internal/svc"
	"modulation/rpc/modulation"
	"runtime/debug"
	"tools/derror"
	"tools/dhttp"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailsLogicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDetailsLogicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailsLogicLogic {
	return &OrderDetailsLogicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderDetailsLogic 二维码订单查询
func (l *OrderDetailsLogicLogic) OrderDetailsLogic(in *modulation.OrderDetailsReq) (result *modulation.OrderDetailsResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			_, err = derror.ResponsePanicErr(nil, r.(string))
			logx.Error(err, "OrderDetailsLogic_recover()", "in", in)
			debug.PrintStack()
		}
	}()
	prams := make(map[string]string, 2)
	prams["OutRequestNo"] = in.OutRequestNo
	prams["PayOutTradeNo"] = in.PayOutTradeNo
	par := ""

	if prams["OutRequestNo"] == "" {
		par = prams["PayOutTradeNo"]
		delete(prams, prams["OutRequestNo"])
	} else if prams["PayOutTradeNo"] == "" {
		par = prams["OutRequestNo"]
		delete(prams, prams["PayOutTradeNo"])
	}

	err = dhttp.ParamsCheck(prams)
	if err != nil {
		obj, err := derror.ResponseErr(nil, err)
		result = InterfaceToStruct(obj)
		return result, err
	}

	rel, err := l.svcCtx.Model.FindOne(l.ctx, par)
	if err != nil {
		logx.Error(err, "redis查询错误", "par", par)
		obj, err := derror.ResponseDBErr(nil, "DB failed to query")
		result = InterfaceToStruct(obj)
		return result, err
	}
	result = &modulation.OrderDetailsResp{
		PayStatus:      rel.PayStatus,
		PayDate:        rel.PayDate,
		PayTime:        rel.PayTime,
		TotalFee:       rel.TotalFee,
		PayCouponFee:   rel.PayCouponFee,
		PayOutTradeNo:  rel.PayOutTradeNo,
		PayErrDesc:     rel.PayErrDesc,
		Uid:            rel.Uid,
		PayType:        rel.PayType,
		PayTypeTradeNo: rel.PayTypeTradeNo,
	}
	obj, err := derror.ResponseSuccess(result)
	result = InterfaceToStruct(obj)
	return result, err
}

// InterfaceToStruct interface转任意类型，每个logic.go中返回类型`*modulation.OrderDetailsResp`记得替换成logic函数的返回类型
func InterfaceToStruct(object interface{}) *modulation.OrderDetailsResp {
	obj := object.(*modulation.OrderDetailsResp)
	return obj
}
