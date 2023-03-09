package logic

import (
	"context"
	"interface/rpc/transform/transformer"

	"interface/api/internal/svc"
	"interface/api/internal/types"

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

func (l *RequestConsumptionLogic) RequestConsumption(req *types.RequestConsumptionRequest) (*types.RequestConsumptionResponse, error) {
	resp, err := l.svcCtx.Transformer.RequestConsumption(l.ctx, &transformer.RequestConsumptionRequest{
		Uid: req.Uid,
	})
	if err != nil {
		return &types.RequestConsumptionResponse{}, err
	}

	return &types.RequestConsumptionResponse{
		Url: resp.Url,
	}, nil

}
