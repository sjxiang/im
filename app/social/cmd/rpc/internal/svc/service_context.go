package svc

import (
	"im/app/social/cmd/rpc/internal/config"
	"im/app/social/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	
	model.FriendModel
	model.FriendApplyModel
	model.GroupModel
	model.GroupApplyModel
	model.GroupMemberModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	
	return &ServiceContext{
		Config:           c,
		FriendModel:      model.NewFriendModel(conn),
		FriendApplyModel: model.NewFriendApplyModel(conn),
		GroupModel:       model.NewGroupModel(conn),
		GroupApplyModel:  model.NewGroupApplyModel(conn),
		GroupMemberModel: model.NewGroupMemberModel(conn),
	}
}
