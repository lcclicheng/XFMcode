package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WecartModel = (*customWecartModel)(nil)

type (
	// WecartModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWecartModel.
	WecartModel interface {
		wecartModel
	}

	customWecartModel struct {
		*defaultWecartModel
	}
)

// NewWecartModel returns a model for the database table.
func NewWecartModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WecartModel {
	return &customWecartModel{
		defaultWecartModel: newWecartModel(conn, c, opts...),
	}
}
