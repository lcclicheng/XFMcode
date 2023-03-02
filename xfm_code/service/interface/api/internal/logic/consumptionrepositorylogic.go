package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConsumptionRepositoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumptionRepositoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConsumptionRepositoryLogic {
	return &ConsumptionRepositoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumptionRepositoryLogic) ConsumptionRepository(req *types.ConsumptionRepositoryRequest) (resp *types.ConsumptionRepositoryResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
