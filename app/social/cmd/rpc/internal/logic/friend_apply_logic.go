package logic

import (
	"context"

	"im/app/social/cmd/rpc/internal/svc"
	"im/app/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendApplyLogic {
	return &FriendApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 1、添加好友
func (l *FriendApplyLogic) FriendApply(in *pb.FriendApplyReq) (*pb.FriendApplyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FriendApplyResp{}, nil
}
