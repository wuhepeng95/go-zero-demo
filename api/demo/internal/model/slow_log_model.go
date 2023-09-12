package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SlowLogModel = (*customSlowLogModel)(nil)

type (
	// SlowLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSlowLogModel.
	SlowLogModel interface {
		slowLogModel
	}

	customSlowLogModel struct {
		*defaultSlowLogModel
	}
)

// NewSlowLogModel returns a model for the database table.
func NewSlowLogModel(conn sqlx.SqlConn) SlowLogModel {
	return &customSlowLogModel{
		defaultSlowLogModel: newSlowLogModel(conn),
	}
}
