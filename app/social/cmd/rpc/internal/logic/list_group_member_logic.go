package logic

import (
	"context"

	"im/app/social/cmd/rpc/internal/svc"
	"im/app/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGroupMemberLogic {
	return &ListGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 6. 列出群成员
func (l *ListGroupMemberLogic) ListGroupMember(in *pb.ListGroupMemberReq) (*pb.ListGroupMemberResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListGroupMemberResp{}, nil
}
