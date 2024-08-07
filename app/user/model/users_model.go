package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		
		ListByNickname(ctx context.Context, nickname string) ([]*Users, error)
		ListByIds(ctx context.Context, ids []string) ([]*Users, error) 
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}


// like 模糊查询
func (m *customUsersModel) ListByNickname(ctx context.Context, nickname string) ([]*Users, error) {
	query := fmt.Sprintf("select %s from %s where `nickname` like ?", usersRows, m.table)
	
	var resp []*Users
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, fmt.Sprint("%", nickname, "%"))	
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// in
func (m *customUsersModel) ListByIds(ctx context.Context, ids []string) ([]*Users, error) {
	query := fmt.Sprintf("select %s from %s where `id` in ('%s')", usersRows, m.table, strings.Join(ids, "','"))

	var resp []*Users
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, ids)	
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}