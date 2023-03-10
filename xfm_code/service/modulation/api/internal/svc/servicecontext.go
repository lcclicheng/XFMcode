package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"modulation/api/internal/config"
	"modulation/rpc/modulationrpc"
)

type ServiceContext struct {
	Config     config.Config
	Modulation modulationrpc.ModulationRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Modulation: modulationrpc.NewModulationRpc(zrpc.MustNewClient(c.Modulation)),
	}
}
