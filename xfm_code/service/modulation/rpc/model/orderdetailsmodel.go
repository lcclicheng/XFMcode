package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderDetailsModel = (*customOrderDetailsModel)(nil)

type (
	// OrderDetailsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderDetailsModel.
	OrderDetailsModel interface {
		orderDetailsModel
	}

	customOrderDetailsModel struct {
		*defaultOrderDetailsModel
	}
)

// NewOrderDetailsModel returns a model for the database table.
func NewOrderDetailsModel(conn sqlx.SqlConn, c cache.CacheConf) OrderDetailsModel {
	return &customOrderDetailsModel{
		defaultOrderDetailsModel: newOrderDetailsModel(conn, c),
	}
}
