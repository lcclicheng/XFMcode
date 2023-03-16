package logic

import (
	"context"
	"wecart/rpc/wecart"

	"wecart/api/internal/svc"
	"wecart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitiateRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitiateRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitiateRequestLogic {
	return &InitiateRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitiateRequestLogic) InitiateRequest(req *types.InitiateRequestReq) (resp *types.InitiateRequestResp, err error) {
	// 手动代码开始
	resp, err := l.svcCtx.WeCart.InitiateRequestLogic(l.ctx, &wecart.InitiateRequestReq{
		Url: req.Url,
	})
	if err != nil {
		return &types.InitiateRequestResp{}, err
	}
	result := &types.Data{}
	for k, v := range resp.InitiateRequestResp {

	}
	return
}
