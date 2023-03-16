package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"wecart/api/internal/config"
	"wecart/rpc/wecartclient"
)

type ServiceContext struct {
	Config config.Config
	WeCart wecartclient.Wecart
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		WeCart: wecartclient.NewWecart(zrpc.MustNewClient(c.WeCart)),
	}
}
