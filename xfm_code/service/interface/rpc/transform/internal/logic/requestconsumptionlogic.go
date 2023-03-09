package logic

import (
	"context"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"

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
	http.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
		f, err := qrcode.Encode("https://www.baidu.com/", qrcode.Highest, 300)
		if err != nil {
			log.Println(err.Error())
			return
		}
		w.Write(f)
	})
	_ = http.ListenAndServe("127.0.0.1:9090", nil)

	return &pb.RequestConsumptionResponse{}, nil
}
