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
	return &types.OrderDetailsResp{
		PayStatus:       resp.PayStatus,
		PayDate:         resp.PayDate,
		PayTime:         resp.PayTime,
		TotalFee:        resp.TotalFee,
		PayCouponFee:    resp.PayCouponFee,
		PayOutTradeNo:   resp.PayOutTradeNo,
		PayErrDesc:      resp.PayErrDesc,
		Uid:             resp.Uid,
		PayType:         resp.PayType,
		PayTypeTradeNo:  resp.PayTypeTradeNo,
		OutRequestNo:    resp.OutRequestNo,
		DimensionalCode: resp.DimensionalCode,
		BarCode:         resp.BarCode,
	}, nil
}
