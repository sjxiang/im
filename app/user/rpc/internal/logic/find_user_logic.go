package logic

import (
	"context"

	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *pb.FindUserReq) (*pb.FindUserResp, error) {
	// 用户搜索，叠加

	var userList []*model.Users
	var err error

	switch {
	case in.Mobile != "":
		user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
		if err == nil {
			userList = append(userList, user)
		}
	case in.Nickname != "":
		userList, err = l.svcCtx.UserModel.ListByNickname(l.ctx, in.Nickname)
	case len(in.Ids) > 0:
		userList, err = l.svcCtx.UserModel.ListByIds(l.ctx, in.Ids)
	}

	if err != nil {
		return nil, err 
	}


	var resp []*pb.User
	copier.Copy(&resp, &userList)

	return &pb.FindUserResp{
		User: resp,
	}, nil
}
