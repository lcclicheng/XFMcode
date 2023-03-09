package logic

import (
	"context"

	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/rpc/transform/internal/svc"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/rpc/transform/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RequestConsumptionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRequestConsumptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestConsumptionLogic {
	return &RequestConsumptionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户申请消费码返回给用户
func (l *RequestConsumptionLogic) RequestConsumption(in *pb.RequestConsumptionRequest) (*pb.RequestConsumptionResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.RequestConsumptionResponse{}, nil
}
