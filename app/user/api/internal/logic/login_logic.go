package logic

import (
	"context"

	"im/app/user/api/internal/svc"
	"im/app/user/api/internal/types"
	"im/app/user/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	
	loginResp, err := l.svcCtx.UserRPC.Login(l.ctx, &pb.LoginReq{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err 
	}

	var reply types.LoginResp
	copier.Copy(&reply, loginResp)

	return &types.LoginResp{
		AccessToken:  reply.AccessToken,
		AccessExpire: reply.AccessExpire,
	}, nil
}
