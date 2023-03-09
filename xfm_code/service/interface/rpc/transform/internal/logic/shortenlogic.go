package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/hash"
	"interface/rpc/transform/internal/svc"
	"interface/rpc/transform/model"
	"interface/rpc/transform/pb"

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
	// 手动代码开始，生成短链接
	key := hash.Md5Hex([]byte(in.Url))[:6]
	object, _ := l.svcCtx.Model.FindOne(l.ctx, key)
	if object != nil {
		return &pb.ShortenResp{
			Shorten: key,
		}, nil
	}

	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{
		Shorten: key,
		Url:     in.Url,
	})
	if err != nil {
		return nil, err
	}

	return &pb.ShortenResp{
		Shorten: key,
	}, nil
	// 手动代码结束
}
