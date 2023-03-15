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
	for k, v := range resp.Data {
		result.OrderDetailsResp[k].PayStatus = v.PayStatus
		result.OrderDetailsResp[k].PayDate = v.PayDate
		result.OrderDetailsResp[k].PayTime = v.PayTime
		result.OrderDetailsResp[k].TotalFee = v.TotalFee
		result.OrderDetailsResp[k].PayCouponFee = v.PayCouponFee
		result.OrderDetailsResp[k].PayOutTradeNo = v.PayOutTradeNo
		result.OrderDetailsResp[k].PayErrDesc = v.PayErrDesc
		result.OrderDetailsResp[k].Uid = v.Uid
		result.OrderDetailsResp[k].PayType = v.PayType
		result.OrderDetailsResp[k].PayTypeTradeNo = v.PayTypeTradeNo
		result.OrderDetailsResp[k].OutRequestNo = v.OutRequestNo
		result.OrderDetailsResp[k].DimensionalCode = v.DimensionalCode
		result.OrderDetailsResp[k].BarCode = v.BarCode
	}
	return result, nil
}
