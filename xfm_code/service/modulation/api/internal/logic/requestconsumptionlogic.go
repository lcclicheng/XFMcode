package logic

import (
	"context"

	"modulation/api/internal/svc"
	"modulation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RequestConsumptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRequestConsumptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestConsumptionLogic {
	return &RequestConsumptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RequestConsumptionLogic) RequestConsumption(req *types.RequestConsumptionRequest) (resp *types.RequestConsumptionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
