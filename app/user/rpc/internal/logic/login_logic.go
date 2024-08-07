package logic

import (
	"context"
	"errors"

	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrMobileNotFound   = errors.New("mobile not found")
var ErrPasswordNotMatch = errors.New("password not match")

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
	case errors.Is(err, model.ErrNotFound):
		return nil, ErrMobileNotFound
	case err != nil:
		return nil, err
	}

	// 密码验证
	if valid, err := util.CheckPassword(user.Password.String, in.Password); err != nil && !valid {
		return nil, ErrPasswordNotMatch
	}

	// 生成 token
	options := util.TokenOptions{
		SecretKey: l.svcCtx.Config.Jwt.SecretKey,
		Duration:  l.svcCtx.Config.Jwt.Duration,
		Fields:    map[string]any{"identity": user.Id},
	}
	token, err := util.GenerateAuth2Token(options)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResp{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
	}, nil
}
