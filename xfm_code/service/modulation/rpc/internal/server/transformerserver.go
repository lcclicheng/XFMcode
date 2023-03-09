// Code generated by goctl. DO NOT EDIT.
// Source: modulation.proto

package server

import (
	"context"

	"rpc/internal/logic"
	"rpc/internal/svc"
	"rpc/pb"
)

type TransformerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTransformerServer
}

func NewTransformerServer(svcCtx *svc.ServiceContext) *TransformerServer {
	return &TransformerServer{
		svcCtx: svcCtx,
	}
}

// 查询消费码状态
func (s *TransformerServer) QueryCodeStatus(ctx context.Context, in *pb.CodeStatusRequest) (*pb.CodeStatusResponse, error) {
	l := logic.NewQueryCodeStatusLogic(ctx, s.svcCtx)
	return l.QueryCodeStatus(in)
}

// 用户申请消费码返回给用户
func (s *TransformerServer) RequestConsumption(ctx context.Context, in *pb.RequestConsumptionRequest) (*pb.RequestConsumptionResponse, error) {
	l := logic.NewRequestConsumptionLogic(ctx, s.svcCtx)
	return l.RequestConsumption(in)
}
