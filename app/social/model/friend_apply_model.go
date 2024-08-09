package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FriendApplyModel = (*customFriendApplyModel)(nil)

type (
	// FriendApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendApplyModel.
	FriendApplyModel interface {
		friendApplyModel
		withSession(session sqlx.Session) FriendApplyModel
	}

	customFriendApplyModel struct {
		*defaultFriendApplyModel
	}
)

// NewFriendApplyModel returns a model for the database table.
func NewFriendApplyModel(conn sqlx.SqlConn) FriendApplyModel {
	return &customFriendApplyModel{
		defaultFriendApplyModel: newFriendApplyModel(conn),
	}
}

func (m *customFriendApplyModel) withSession(session sqlx.Session) FriendApplyModel {
	return NewFriendApplyModel(sqlx.NewSqlConnFromSession(session))
}
