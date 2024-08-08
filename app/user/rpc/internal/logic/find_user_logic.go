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
	
	// 用户搜索（用户名、wxid、手机号），结果可以为空

	var items []*model.User
	var err error

	if in.Nickname != "" {
		items, err = l.svcCtx.UserModel.ListByNickname(l.ctx, in.Nickname)
	} else if len(in.Ids) > 0 {
		items, err = l.svcCtx.UserModel.ListByIds(l.ctx, in.Ids)
	} 

	if err != nil {
		return nil, err 
	}

	var resp []*pb.User
	copier.Copy(&resp, &items)

	return &pb.FindUserResp{
		Users: resp,
	}, nil
}

