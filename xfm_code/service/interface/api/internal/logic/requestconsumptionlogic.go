package logic

import (
	"context"

	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/api/internal/svc"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/api/internal/types"

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
