package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GroupApplyModel = (*customGroupApplyModel)(nil)

type (
	// GroupApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupApplyModel.
	GroupApplyModel interface {
		groupApplyModel
		withSession(session sqlx.Session) GroupApplyModel
	}

	customGroupApplyModel struct {
		*defaultGroupApplyModel
	}
)

// NewGroupApplyModel returns a model for the database table.
func NewGroupApplyModel(conn sqlx.SqlConn) GroupApplyModel {
	return &customGroupApplyModel{
		defaultGroupApplyModel: newGroupApplyModel(conn),
	}
}

func (m *customGroupApplyModel) withSession(session sqlx.Session) GroupApplyModel {
	return NewGroupApplyModel(sqlx.NewSqlConnFromSession(session))
}
