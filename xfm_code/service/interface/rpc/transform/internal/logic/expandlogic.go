package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"interface/rpc/transform/internal/svc"
	"interface/rpc/transform/pb"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *pb.ExpandReq) (*pb.ExpandResp, error) {
	// 手动代码开始
	res, err := l.svcCtx.Model.FindOne(l.ctx, in.Shorten)
	if err != nil {
		return nil, err
	}

	return &pb.ExpandResp{
		Url: res.Url,
	}, nil
	// 手动代码结束
}
