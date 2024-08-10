package model

import (
	"fmt"
	"context"
	"strings"
	
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel

		ListByNickname(ctx context.Context, nickname string) ([]*User, error) 
		ListByIds(ctx context.Context, ids []string) ([]*User, error) 
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

// like 模糊查询
func (m *customUserModel) ListByNickname(ctx context.Context, nickname string) ([]*User, error) {
	query := fmt.Sprintf("select %s from %s where `nickname` like %s", userRows, m.table, "'%"+ nickname+ "%'")
	
	var resp []*User
	err := m.conn.QueryRowsCtx(ctx, &resp, query)	
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// in
func (m *customUserModel) ListByIds(ctx context.Context, ids []string) ([]*User, error) {
	
	query := fmt.Sprintf("select %s from `%s` where `id` in ('%s')", userRows, "users", strings.Join(ids, "','"))

	var resp []*User
	err := m.conn.QueryRowsCtx(ctx, &resp, query)	
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

/*

select 
	`id`,`nickname`,`mobile`,`password`,`sex`,`status`,`intro`,`avatar`,`created_at`,`updated_at` 
from 
	`users` 
where 
	`id` in ('602ebd7f-d8fb-4a3f-9f0b-73ec0656edd0','eb76279e-1adf-4bd6-9fa5-9033a1ebe982');

	
select 
	`id`,`nickname`,`mobile`,`password`,`sex`,`status`,`intro`,`avatar`,`created_at`,`updated_at` 
from 
	`users` 
where 
	`id` = '602ebd7f-d8fb-4a3f-9f0b-73ec0656edd0' l
imit 1;


select 
	`id`,`nickname`,`mobile`,`password`,`sex`,`status`,`intro`,`avatar`,`created_at`,`updated_at` 
from 
	`users` 
where 
	`nickname` like '%jisoo%';

 */

