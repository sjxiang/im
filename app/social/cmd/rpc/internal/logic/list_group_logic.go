package logic

import (
	"context"

	"im/app/social/cmd/rpc/internal/svc"
	"im/app/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGroupLogic {
	return &ListGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 5. 列出群组
func (l *ListGroupLogic) ListGroup(in *pb.ListGroupReq) (*pb.ListGroupResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListGroupResp{}, nil
}
