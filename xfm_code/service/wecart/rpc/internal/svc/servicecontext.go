package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wecart/rpc/internal/config"
	"wecart/rpc/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.WecartModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewWecartModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
