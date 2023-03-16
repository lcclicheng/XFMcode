package logic

import (
	"context"

	"wecart/rpc/internal/svc"
	"wecart/rpc/wecart"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitiateRequestLogicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitiateRequestLogicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitiateRequestLogicLogic {
	return &InitiateRequestLogicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitiateRequestLogicLogic) InitiateRequestLogic(in *wecart.InitiateRequestReq) (*wecart.InitiateRequestResp, error) {
	// 手动代码开始
	res, err := l.svcCtx.Model.FindOne(l.ctx, in.Url)
	if err != nil {
		return nil, err
	}

	return &wecart.InitiateRequestResp{
		Data: res.Url,
	}, nil
	// 手动代码结束
	return &wecart.InitiateRequestResp{}, nil
}
