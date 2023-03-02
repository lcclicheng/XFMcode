package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentAgreementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentAgreementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentAgreementLogic {
	return &PaymentAgreementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentAgreementLogic) PaymentAgreement(req *types.PaymentAgreementRequest) (resp *types.PaymentAgreementResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
