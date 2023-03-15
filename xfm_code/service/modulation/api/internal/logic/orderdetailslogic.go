package logic

import (
	"context"
	"modulation/rpc/modulation"

	"modulation/api/internal/svc"
	"modulation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailsLogic {
	return &OrderDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailsLogic) OrderDetails(req *types.OrderDetailsReq) (*types.OrderDetailsResp, error) {
	resp, err := l.svcCtx.Modulation.OrderDetailsLogic(l.ctx, &modulation.OrderDetailsReq{
		OutRequestNo:  req.OutRequestNo,
		PayOutTradeNo: req.PayOutTradeNo,
	})
	if err != nil {
		return &types.OrderDetailsResp{}, err
	}
	result := &types.OrderDetailsResp{}
	for _, v1 := range resp.Data {
		for _, v2 := range result.OrderDetailsResp {
			v2.PayStatus = v1.PayStatus
			v2.PayDate = v1.PayDate
			v2.PayTime = v1.PayTime
			v2.TotalFee = v1.TotalFee
			v2.PayCouponFee = v1.PayCouponFee
			v2.PayOutTradeNo = v1.PayOutTradeNo
			v2.PayErrDesc = v1.PayErrDesc
			v2.Uid = v1.Uid
			v2.PayType = v1.PayType
			v2.PayTypeTradeNo = v1.PayTypeTradeNo
			v2.OutRequestNo = v1.OutRequestNo
			v2.DimensionalCode = v1.DimensionalCode
			v2.BarCode = v1.BarCode
		}
	}
	return result, nil
}
