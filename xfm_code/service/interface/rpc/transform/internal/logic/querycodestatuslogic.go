package logic

import (
	"context"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/rpc/transform/internal/svc"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/rpc/transform/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCodeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCodeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCodeStatusLogic {
	return &QueryCodeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询消费码状态
func (l *QueryCodeStatusLogic) QueryCodeStatus(in *pb.CodeStatusRequest) (*pb.CodeStatusResponse, error) {
	//_, err := l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{
	//	Shorten: key,
	//	Url:     in.Url,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &pb.ShortenResp{
	//	Shorten: key,
	//}, nil

	return &pb.CodeStatusResponse{Status: "已领取个人消费码"}, nil
}
