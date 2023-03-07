package svc

import (
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/api/internal/config"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/rpc/transform/transformer"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
	}
}
