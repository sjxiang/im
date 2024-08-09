package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GroupMemberModel = (*customGroupMemberModel)(nil)

type (
	// GroupMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupMemberModel.
	GroupMemberModel interface {
		groupMemberModel
		withSession(session sqlx.Session) GroupMemberModel
	}

	customGroupMemberModel struct {
		*defaultGroupMemberModel
	}
)

// NewGroupMemberModel returns a model for the database table.
func NewGroupMemberModel(conn sqlx.SqlConn) GroupMemberModel {
	return &customGroupMemberModel{
		defaultGroupMemberModel: newGroupMemberModel(conn),
	}
}

func (m *customGroupMemberModel) withSession(session sqlx.Session) GroupMemberModel {
	return NewGroupMemberModel(sqlx.NewSqlConnFromSession(session))
}
