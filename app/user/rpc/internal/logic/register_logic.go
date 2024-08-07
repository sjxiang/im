package logic

import (
	"context"

	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	
	// 验证是否存在用户，即手机号码

	// 对密码加密（也有可能是 mobile + verify code 搭配）
	hashedPassword, err := util.HashPassword(in.Password)
	if err != nil {
		return nil, err
	}

	// 新增用户
	
	_ = hashedPassword 

	// 生成 token

	return &pb.RegisterResp{}, nil
}
