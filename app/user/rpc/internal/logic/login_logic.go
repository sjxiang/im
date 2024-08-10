package logic

import (
	"context"
	"fmt"

	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"im/pkg/util"
	"im/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrPhoneNotFound    = xerr.NewCodeMsg(xerr.SERVER_COMMON_ERROR, "该手机号码未注册")
var ErrPasswordNotMatch = xerr.NewCodeMsg(xerr.SERVER_COMMON_ERROR, "输入密码不匹配")
var ErrTokenGenFailed   = xerr.NewCodeMsg(xerr.TOKEN_GENERATE_Failed_ERROR, "生成 token 失败")

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
	user, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone) 
	switch {
	case err == model.ErrNotFound:
		return nil, errors.WithStack(ErrPhoneNotFound)
	case err != nil:
		return nil, errors.Wrapf(xerr.NewDBErr(), fmt.Sprintf("find user by mobile, err %v, req %s", err, in.Phone))
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
		return nil, errors.Wrapf(ErrTokenGenFailed, fmt.Sprintf("generate token failed, err %v", err))
	}

	return &pb.LoginResp{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
	}, nil
}
