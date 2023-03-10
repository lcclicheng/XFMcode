package logic

import (
	"context"

	"modulation/rpc/internal/svc"
	"modulation/rpc/modulation"

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

// 二维码订单查询
func (l *OrderDetailsLogicLogic) OrderDetailsLogic(in *modulation.OrderDetailsReq) (*modulation.OrderDetailsResp, error) {
	// todo: add your logic here and delete this line

	return &modulation.OrderDetailsResp{}, nil
}
