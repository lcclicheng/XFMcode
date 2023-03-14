package logic

import (
	"context"
	"encoding/base64"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"modulation/rpc/internal/svc"
	"modulation/rpc/modulation"
	"os"

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
func (l *RequestConsumptionLogic) RequestConsumption(in *modulation.RequestConsumptionRequest) (*modulation.RequestConsumptionResponse, error) {
	//生成二维码
	qrCode, _ := qr.Encode("https://blog.csdn.net/weixin_36162966", qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)
	file, _ := os.Create("qr1.png")
	defer file.Close()
	err := png.Encode(file, qrCode)
	if err != nil {
		return nil, err
	}
	f, err := os.Open("qr1.png")
	if err != nil {
		panic(err)
	}
	_, _, err = image.Decode(f)
	if err != nil {
		panic(err)
	}
	srcByte, err := ioutil.ReadFile(`qr1.png`)
	if err != nil {
		log.Fatal(err)
	}
	res := base64.StdEncoding.EncodeToString(srcByte)
	//生成条形码
	cs, _ := code128.Encode("123456")
	files, _ := os.Create("qr2.png")
	defer files.Close()
	qrCodes, _ := barcode.Scale(cs, 350, 70)
	err = png.Encode(files, qrCodes)
	if err != nil {
		return nil, err
	}
	f, err = os.Open("qr1.png")
	if err != nil {
		panic(err)
	}
	_, _, err = image.Decode(f)
	if err != nil {
		panic(err)
	}
	srcBytes, err := ioutil.ReadFile(`qr2.png`)
	if err != nil {
		log.Fatal(err)
	}
	result := base64.StdEncoding.EncodeToString(srcBytes)
	return &modulation.RequestConsumptionResponse{DimensionalCode: res, BarCode: result}, nil
}
