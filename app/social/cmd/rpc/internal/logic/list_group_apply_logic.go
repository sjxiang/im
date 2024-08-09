package logic

import (
	"context"

	"im/app/social/cmd/rpc/internal/svc"
	"im/app/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListGroupApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListGroupApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGroupApplyLogic {
	return &ListGroupApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 3、列出群组申请
func (l *ListGroupApplyLogic) ListGroupApply(in *pb.ListGroupApplyReq) (*pb.ListGroupApplyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListGroupApplyResp{}, nil
}
