package logic

import (
	"context"

	"modulation/rpc/internal/svc"
	"modulation/rpc/modulation"

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
func (l *QueryCodeStatusLogic) QueryCodeStatus(in *modulation.CodeStatusRequest) (*modulation.CodeStatusResponse, error) {
	res, err := l.svcCtx.Model.FindOne(l.ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	if len(res.DimensionalCode) != 0 {
		return &modulation.CodeStatusResponse{Status: "已领取个人消费码"}, nil
	} else {
		return &modulation.CodeStatusResponse{Status: "未领取个人消费码"}, nil
	}
	return nil, nil
}
