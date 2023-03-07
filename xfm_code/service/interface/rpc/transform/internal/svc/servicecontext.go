package svc

import "github.com/lcclicheng/XFMcode/xfm_code/service/interface/rpc/transform/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
