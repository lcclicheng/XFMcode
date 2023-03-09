package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

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
	// todo: add your logic here and delete this line

	return &pb.CodeStatusResponse{}, nil
}
