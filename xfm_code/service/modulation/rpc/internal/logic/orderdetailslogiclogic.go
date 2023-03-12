package logic

import (
	"context"
	"fmt"
	"modulation/rpc/internal/svc"
	"modulation/rpc/modulation"
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
func (l *OrderDetailsLogicLogic) OrderDetailsLogic(in *modulation.OrderDetailsReq) (*modulation.OrderDetailsResp, error) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
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

	err := dhttp.ParamsCheck(prams)
	if err != nil {
		fmt.Printf("Regular check failed err:%s", err)
		return nil, err
	}

	rel, err := l.svcCtx.Model.FindOne(l.ctx, par)
	if err != nil {
		fmt.Printf("Query DB Failed err:%s", err)
		return nil, err
	}

	return &modulation.OrderDetailsResp{
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
	}, nil
}
