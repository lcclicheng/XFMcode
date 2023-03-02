package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

//--从支付方获取支付签约请求结果--

type GetPaymentRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPaymentRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentRequestLogic {
	return &GetPaymentRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPaymentRequestLogic) GetPaymentRequest(req *types.GetPaymentRequestReq) (resp *types.GetPaymentRequestResponse, err error) {

	return
}
