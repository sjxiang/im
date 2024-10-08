// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupMemberFieldNames          = builder.RawFieldNames(&GroupMember{})
	groupMemberRows                = strings.Join(groupMemberFieldNames, ",")
	groupMemberRowsExpectAutoSet   = strings.Join(stringx.Remove(groupMemberFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupMemberRowsWithPlaceHolder = strings.Join(stringx.Remove(groupMemberFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	groupMemberModel interface {
		Insert(ctx context.Context, data *GroupMember) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*GroupMember, error)
		Update(ctx context.Context, data *GroupMember) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultGroupMemberModel struct {
		conn  sqlx.SqlConn
		table string
	}

	GroupMember struct {
		Id          uint64         `db:"id"`
		GroupId     string         `db:"group_id"`
		UserId      string         `db:"user_id"`
		RoleLevel   int64          `db:"role_level"` // 群聊用户等级
		JoinAt      sql.NullTime   `db:"join_at"`
		JoinSource  sql.NullInt64  `db:"join_source"`
		InviterUid  sql.NullString `db:"inviter_uid"`  // 邀请人
		OperatorUid sql.NullString `db:"operator_uid"` // 审核员
	}
)

func newGroupMemberModel(conn sqlx.SqlConn) *defaultGroupMemberModel {
	return &defaultGroupMemberModel{
		conn:  conn,
		table: "`group_member`",
	}
}

func (m *defaultGroupMemberModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGroupMemberModel) FindOne(ctx context.Context, id uint64) (*GroupMember, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupMemberRows, m.table)
	var resp GroupMember
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

func (m *defaultGroupMemberModel) Insert(ctx context.Context, data *GroupMember) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, groupMemberRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.GroupId, data.UserId, data.RoleLevel, data.JoinAt, data.JoinSource, data.InviterUid, data.OperatorUid)
	return ret, err
}

func (m *defaultGroupMemberModel) Update(ctx context.Context, data *GroupMember) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupMemberRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.GroupId, data.UserId, data.RoleLevel, data.JoinAt, data.JoinSource, data.InviterUid, data.OperatorUid, data.Id)
	return err
}

func (m *defaultGroupMemberModel) tableName() string {
	return m.table
}
