package logic

import (
	"context"
	"fmt"

	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"im/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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
	
	// 用户搜索（用户名、wxid、手机号），结果可以为空

	var items []*model.User
	var err error

	if in.Nickname != "" {
		items, err = l.svcCtx.UserModel.ListByNickname(l.ctx, in.Nickname)
	} else if len(in.Ids) > 0 {
		items, err = l.svcCtx.UserModel.ListByIds(l.ctx, in.Ids)
	} else if in.Phone != "" {
		item, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
		if err == nil {
			items = append(items, item)
		}
	}

	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), fmt.Sprintf("find user by phone or id, err %v", err))
	}

	var resp []*pb.User
	copier.Copy(&resp, &items)
	
	l.Logger.Infow("[🚀查找用户]", logx.Field("用户数据", &resp))

	return &pb.FindUserResp{
		Users: resp,
	}, nil
}

