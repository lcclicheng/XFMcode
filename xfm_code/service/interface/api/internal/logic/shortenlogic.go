package logic

import (
	"api/internal/common"
	"context"
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	qr, err := qrcode.New("https://www.flysnow.org/", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{R: 50, G: 205, B: 50, A: 255}
		qr.ForegroundColor = color.White
		if err = qr.WriteFile(256, "./blog_qrcode.png"); err != nil {
			return nil, common.NewCodeError(1003, "生成二维码失败")
		}
	}
	resp.Shorten = "测试"
	return resp, nil
}
