package logic

import (
	"context"
	"wecart/rpc/wecart"

	"wecart/api/internal/svc"
	"wecart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceAnOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceAnOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceAnOrderLogic {
	return &PlaceAnOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceAnOrderLogic) PlaceAnOrder(req *types.PlaceAnOrderReq) (*types.PlaceAnOrderResp, error) {
	resp, err := l.svcCtx.WeCart.PlaceAnOrderLogic(l.ctx, &wecart.PlaceAnOrderReq{
		Appid:       req.Appid,
		Mchid:       req.Mchid,
		Description: req.Description,
		OutTradeNo:  req.OutTradeNo,
		Attach:      req.Attach,
		NotifyUrl:   req.NotifyUrl,
		Openid:      req.Openid,
	})
	if err != nil {
		return &types.PlaceAnOrderResp{}, err
	}
	result := &types.PlaceAnOrderResp{}
	for k, v := range resp.Data {
		result.PlaceAnOrderResp[k].PrepayId = v.PrepayId
		result.PlaceAnOrderResp[k].Appid = v.Appid
		result.PlaceAnOrderResp[k].TimeStamp = v.TimeStamp
		result.PlaceAnOrderResp[k].NonceStr = v.NonceStr
		result.PlaceAnOrderResp[k].Package = v.Package
		result.PlaceAnOrderResp[k].SignType = v.SignType
		result.PlaceAnOrderResp[k].PaySign = v.PaySign
	}

	return result, nil
}
