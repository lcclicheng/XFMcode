package logic

import (
	"api/internal/common/errorx"
	"api/internal/svc"
	"api/internal/types"
	"context"
	"github.com/skip2/go-qrcode"
	"github.com/zeromicro/go-zero/core/logx"
	"image/color"
	"log"
)

//--用户申请消费码返回给用户--

type RequestConsumptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRequestConsumptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestConsumptionLogic {
	return &RequestConsumptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RequestConsumptionLogic) RequestConsumption(req *types.RequestConsumptionRequest) (resp *types.RequestConsumptionResponse, err error) {
	qr, err := qrcode.New("https://www.flysnow.org/", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{R: 50, G: 205, B: 50, A: 255}
		qr.ForegroundColor = color.White
		if err = qr.WriteFile(256, "./blog_qrcode.png"); err != nil {
			return nil, errorx.NewCodeError(1003, "生成二维码失败")
		}
	}
	return
}
