package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"modulation/rpc/internal/config"
	"modulation/rpc/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.OrderDetailsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewOrderDetailsModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
