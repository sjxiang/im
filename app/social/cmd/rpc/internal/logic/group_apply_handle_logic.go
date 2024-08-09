package logic

import (
	"context"

	"im/app/social/cmd/rpc/internal/svc"
	"im/app/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupApplyHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupApplyHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupApplyHandleLogic {
	return &GroupApplyHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 4. 处理加群申请
func (l *GroupApplyHandleLogic) GroupApplyHandle(in *pb.GroupApplyHandleReq) (*pb.GroupApplyHandleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GroupApplyHandleResp{}, nil
}
