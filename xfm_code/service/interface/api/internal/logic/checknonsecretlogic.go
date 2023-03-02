package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckNonSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckNonSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckNonSecretLogic {
	return &CheckNonSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckNonSecretLogic) CheckNonSecret(req *types.CheckNonSecretRequest) (resp *types.CheckNonSecretResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
