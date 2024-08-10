package logic

import (
	"context"
	"fmt"

	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"im/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)


var ErrUserNotFound = xerr.NewCodeMsg(xerr.SERVER_COMMON_ERROR, "不存在该用户")

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
	case err == model.ErrNotFound:
		return nil, errors.WithStack(ErrUserNotFound)
	case err != nil:
		return nil, errors.Wrapf(xerr.NewDBErr(), fmt.Sprintf("find user by id, err %v, param %s", err, in.Id))
	}

	var resp pb.User
	copier.Copy(&resp, user)  // 第一个参数是要设置的对象，第二个参数是数据的来源

	l.Logger.Infow("[🚀获取用户信息]", logx.Field("用户数据", &resp))

	return &pb.GetUserInfoResp{
		User: &resp,
	}, nil
}
