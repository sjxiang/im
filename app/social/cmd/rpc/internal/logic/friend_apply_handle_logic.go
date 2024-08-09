package logic

import (
	"context"

	"im/app/social/cmd/rpc/internal/svc"
	"im/app/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendApplyHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendApplyHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendApplyHandleLogic {
	return &FriendApplyHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 2、处理好友申请
func (l *FriendApplyHandleLogic) FriendApplyHandle(in *pb.FriendApplyHandleReq) (*pb.FriendApplyHandleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FriendApplyHandleResp{}, nil
}
