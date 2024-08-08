package logic

import (
	"context"

	"im/app/user/api/internal/svc"
	"im/app/user/api/internal/types"
	"im/app/user/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {

	registerResp, err := l.svcCtx.UserRPC.Register(l.ctx, &pb.RegisterReq{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Password: req.Password,
		Avatar:   req.Avatar,
		Sex:      req.Sex,
	})
	if err != nil {
		return nil, err 
	}

	var reply types.RegisterResp
	copier.Copy(&reply, registerResp)

	return &types.RegisterResp{
		UserId: reply.UserId,
	}, nil
}
