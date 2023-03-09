package logic

import (
	"context"
	"interface/rpc/transform/transformer"

	"interface/api/internal/svc"
	"interface/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCodeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCodeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCodeStatusLogic {
	return &QueryCodeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCodeStatusLogic) QueryCodeStatus(req *types.CodeStatusRequest) (*types.CodeStatusResponse, error) {
	resp, err := l.svcCtx.Transformer.QueryCodeStatus(l.ctx, &transformer.CodeStatusRequest{
		Uid: req.Uid,
	})
	if err != nil {
		return &types.CodeStatusResponse{}, err
	}

	return &types.CodeStatusResponse{
		Status: resp.Status,
	}, nil
}
