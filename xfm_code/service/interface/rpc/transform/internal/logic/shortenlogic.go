package logic

import (
	"context"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/rpc/transform/internal/svc"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/rpc/transform/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *pb.ShortenReq) (*pb.ShortenResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ShortenResp{}, nil
}
