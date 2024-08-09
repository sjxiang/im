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
	groupFieldNames          = builder.RawFieldNames(&Group{})
	groupRows                = strings.Join(groupFieldNames, ",")
	groupRowsExpectAutoSet   = strings.Join(stringx.Remove(groupFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupRowsWithPlaceHolder = strings.Join(stringx.Remove(groupFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	groupModel interface {
		Insert(ctx context.Context, data *Group) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Group, error)
		Update(ctx context.Context, data *Group) error
		Delete(ctx context.Context, id string) error
	}

	defaultGroupModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Group struct {
		Id              string         `db:"id"`
		Name            string         `db:"name"`        // 群名
		Icon            string         `db:"icon"`        // 群图标
		Status          sql.NullInt64  `db:"status"`      // 是否
		CreatorUid      string         `db:"creator_uid"` // 创建群组的用户id
		GroupType       int64          `db:"group_type"`
		IsVerify        bool           `db:"is_verify"`    // 是否开启验证
		Notification    sql.NullString `db:"notification"` // 群公告
		NotificationUid sql.NullString `db:"notification_uid"`
		CreatedAt       time.Time      `db:"created_at"` // 添加时间
		UpdatedAt       time.Time      `db:"updated_at"` // 更新时间
	}
)

func newGroupModel(conn sqlx.SqlConn) *defaultGroupModel {
	return &defaultGroupModel{
		conn:  conn,
		table: "`group`",
	}
}

func (m *defaultGroupModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGroupModel) FindOne(ctx context.Context, id string) (*Group, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupRows, m.table)
	var resp Group
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

func (m *defaultGroupModel) Insert(ctx context.Context, data *Group) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, groupRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.Icon, data.Status, data.CreatorUid, data.GroupType, data.IsVerify, data.Notification, data.NotificationUid)
	return ret, err
}

func (m *defaultGroupModel) Update(ctx context.Context, data *Group) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Icon, data.Status, data.CreatorUid, data.GroupType, data.IsVerify, data.Notification, data.NotificationUid, data.Id)
	return err
}

func (m *defaultGroupModel) tableName() string {
	return m.table
}
