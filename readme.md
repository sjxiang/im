

# im

>



package logic

import (
	"context"
	"database/sql"
	"errors"

	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrMobileAlreadyExists = errors.New("mobile already exist")
var ErrMobileNotFound      = errors.New("mobile not found")

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
	
	// 验证用户是否注册，根据手机号码验证
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile) 
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err  
	}
	if user != nil {
		return nil, ErrMobileAlreadyExists
	}

	// 定义用户数据
	user = &model.Users{
		Id: 1,
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Mobile:   in.Mobile,
		Sex:      in.Sex,
	}

	// 对密码加密（有可能不需要密码，而是手机验证码直接注册）
	if len(in.Password) > 0 {
		hashedPasword, err := util.HashPassword(in.Password)
		if err != nil {
			return nil, err 
		}

		user.Password = sql.NullString{
			String: hashedPasword,
			Valid:  true,
		}
	}

	// 新增用户
	if _, err := l.svcCtx.UserModel.Insert(l.ctx, user); err != nil {
		return nil, err 
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

	return &pb.RegisterResp{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
	}, nil
}
