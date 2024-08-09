package logic

import (
	"context"

	"im/app/social/cmd/rpc/internal/svc"
	"im/app/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFriendApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFriendApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFriendApplyLogic {
	return &ListFriendApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 3、列出好友申请
func (l *ListFriendApplyLogic) ListFriendApply(in *pb.FriendApplyListReq) (*pb.FriendApplyListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FriendApplyListResp{}, nil
}
