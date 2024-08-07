package logic

import (
	"context"

	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *pb.FindUserReq) (*pb.FindUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FindUserResp{}, nil
}
