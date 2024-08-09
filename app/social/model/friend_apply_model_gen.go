// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	friendApplyFieldNames          = builder.RawFieldNames(&FriendApply{})
	friendApplyRows                = strings.Join(friendApplyFieldNames, ",")
	friendApplyRowsExpectAutoSet   = strings.Join(stringx.Remove(friendApplyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	friendApplyRowsWithPlaceHolder = strings.Join(stringx.Remove(friendApplyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	friendApplyModel interface {
		Insert(ctx context.Context, data *FriendApply) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*FriendApply, error)
		Update(ctx context.Context, data *FriendApply) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultFriendApplyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	FriendApply struct {
		Id           uint64       `db:"id"`            // 自增id
		UserId       string       `db:"user_id"`       // 用户id
		ApplyUid     string       `db:"apply_uid"`     // 申请用户id
		ApplyMsg     string       `db:"apply_msg"`     // 申请信息
		ApplyAt      time.Time    `db:"apply_at"`      // 申请时间
		HandleResult int64        `db:"handle_result"` // 处理结果
		HandleMsg    string       `db:"handle_msg"`    // 处理回复
		HandleAt     sql.NullTime `db:"handle_at"`     // 添加时间
	}
)

func newFriendApplyModel(conn sqlx.SqlConn) *defaultFriendApplyModel {
	return &defaultFriendApplyModel{
		conn:  conn,
		table: "`friend_apply`",
	}
}

func (m *defaultFriendApplyModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultFriendApplyModel) FindOne(ctx context.Context, id uint64) (*FriendApply, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendApplyRows, m.table)
	var resp FriendApply
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFriendApplyModel) Insert(ctx context.Context, data *FriendApply) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, friendApplyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ApplyUid, data.ApplyMsg, data.ApplyAt, data.HandleResult, data.HandleMsg, data.HandleAt)
	return ret, err
}

func (m *defaultFriendApplyModel) Update(ctx context.Context, data *FriendApply) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, friendApplyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ApplyUid, data.ApplyMsg, data.ApplyAt, data.HandleResult, data.HandleMsg, data.HandleAt, data.Id)
	return err
}

func (m *defaultFriendApplyModel) tableName() string {
	return m.table
}
