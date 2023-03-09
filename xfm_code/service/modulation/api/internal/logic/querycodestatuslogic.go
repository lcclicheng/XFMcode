package logic

import (
	"context"

	"modulation/api/internal/svc"
	"modulation/api/internal/types"

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

func (l *QueryCodeStatusLogic) QueryCodeStatus(req *types.CodeStatusRequest) (resp *types.CodeStatusResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
