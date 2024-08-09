package logic

import (
	"context"

	"im/app/social/cmd/rpc/internal/svc"
	"im/app/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupApplyLogic {
	return &GroupApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 2. 添加群众
func (l *GroupApplyLogic) GroupApply(in *pb.GroupApplyReq) (*pb.GroupApplyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GroupApplyResp{}, nil
}
