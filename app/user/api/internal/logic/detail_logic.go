package logic

import (
	"context"
	"errors"

	"im/app/user/api/internal/svc"
	"im/app/user/api/internal/types"
	"im/app/user/rpc/pb"
	"im/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail() (resp *types.UserInfoResp, err error) {
	
	uid := util.GetUidFromCtx(l.ctx)
	if uid == "" {
		return nil, errors.New("uid is empty")
	}
	l.Logger.Info("用户id", uid)

	getUserInfoResp, err := l.svcCtx.UserRPC.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: uid,
	})
	if err != nil {
		return nil, err 
	}

	var reply types.User
	copier.Copy(&reply, getUserInfoResp.User)

	return &types.UserInfoResp{
		UserInfo: reply,
	}, nil
}