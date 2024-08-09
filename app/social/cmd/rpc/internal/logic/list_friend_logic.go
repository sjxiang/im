package logic

import (
	"context"

	"im/app/social/cmd/rpc/internal/svc"
	"im/app/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFriendLogic {
	return &ListFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 4、列出好友
func (l *ListFriendLogic) ListFriend(in *pb.ListFriendReq) (*pb.ListFriendResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListFriendResp{}, nil
}
