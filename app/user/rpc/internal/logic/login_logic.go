package logic

import (
	"context"

	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"im/pkg/util"
	"im/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrMobileNotFound   = xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "该手机号码未注册")
var ErrPasswordNotMatch = xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "输入密码不匹配")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	
	// 验证用户是否注册，根据手机号码验证
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile) 
	switch {
	case err == model.ErrNotFound:
		return nil, errors.WithStack(ErrMobileNotFound)
	case err != nil:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find user by mobile err #{err}, req #{in.Mobile}")
	}

	// 密码验证
	if valid, err := util.CheckPassword(user.Password.String, in.Password); err != nil && !valid {
		return nil, errors.WithStack(ErrPasswordNotMatch)
	}

	// 生成 token
	options := util.TokenOptions{
		SecretKey: l.svcCtx.Config.Jwt.SecretKey,
		Duration:  l.svcCtx.Config.Jwt.Duration,
		Fields:    map[string]any{"identity": user.Id},
	}
	token, err := util.GenerateAuth2Token(options)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "generate token err %v", err)
	}

	return &pb.LoginResp{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
	}, nil
}
