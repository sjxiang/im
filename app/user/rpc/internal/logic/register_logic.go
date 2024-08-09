package logic

import (
	"context"
	"database/sql"
	
	"im/app/user/model"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"im/pkg/util"
	"im/pkg/xerr"
	
	"github.com/pkg/errors"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrMobileAlreadyExists = xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "手机号码已经注册")
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
	
	// 验证用户是否注册，根据手机号码验证
	record, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone) 
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "find user by mobile, err #{err}, param #{in.Mobile}")
	case record != nil:
		return nil, errors.WithStack(ErrMobileAlreadyExists)
	}

	// 定义用户数据
	newUser := &model.User{
		Id:       uuid.New().String(),
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Sex:      in.Sex,
	}

	// 对密码加密（有可能不需要密码，而是手机验证码直接注册）
	if len(in.Password) > 0 {
		hashedPasword, err := util.HashPassword(in.Password)
		if err != nil {
			return nil, errors.Wrap(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "hash password, err #{err}")
		}

		newUser.Password = sql.NullString{
			String: hashedPasword,
			Valid:  true,
		}
	}

	// 新增用户
	_, err = l.svcCtx.UserModel.Insert(l.ctx, newUser)
	if err != nil {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "insert user, err #{err}, param #{newUser}")
	}
	
	l.Logger.Infow("[🚀注册]", logx.Field("用户数据", newUser.Id))

	return &pb.RegisterResp{
		Id: newUser.Id,
	}, nil
}
