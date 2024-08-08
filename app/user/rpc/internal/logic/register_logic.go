package logic

import (
	"context"
	"database/sql"
	"errors"

	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"im/pkg/util"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrMobileAlreadyExists = errors.New("mobile already exist")  // 改手机号码已经注册

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
	record, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile) 
	switch {
	case err != nil && !errors.Is(err, model.ErrNotFound):
		return nil, err 
	case record != nil:
		return nil, ErrMobileAlreadyExists
	}

	// 定义用户数据
	newUser := &model.User{
		Id:       uuid.New().String(),
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

		newUser.Password = sql.NullString{
			String: hashedPasword,
			Valid:  true,
		}
	}

	// 新增用户
	_, err = l.svcCtx.UserModel.Insert(l.ctx, newUser)
	if err != nil {
		return nil, err 
	}
	

	return &pb.RegisterResp{
		Id: newUser.Id,
	}, nil
}
