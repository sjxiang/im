package logic

import (
	"context"
	"errors"

	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)


var ErrUserNotFound = errors.New("user not found")  // 不存在该用户

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.GetId())
	switch {
	case errors.Is(err, model.ErrNotFound):
		return nil, ErrUserNotFound
	case err != nil:
		return nil, err 
	}

	var resp pb.User
	copier.Copy(&resp, user)  // 第一个参数是要设置的对象，第二个参数是数据的来源

	return &pb.GetUserInfoResp{
		User: &resp,
	}, nil
}
