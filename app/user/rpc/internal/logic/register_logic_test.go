package logic

import (
	"context"
	"testing"

	"im/app/user/rpc/internal/config"
	"im/app/user/rpc/internal/svc"
	"im/app/user/rpc/pb"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/conf"
)

var svcCtx *svc.ServiceContext

func init() {
	var cfg config.Config
	conf.MustLoad(filepath.Join("../../etc/user.yaml"), &cfg)
	svcCtx = svc.NewServiceContext(cfg)
}

func Test_register_logic(t *testing.T) {
	var tests = []struct{
		name string
		args *pb.RegisterReq
		expect bool
	}{
		{
			name: "注册成功",
			args: &pb.RegisterReq{
				Nickname: "jisoo",
				Mobile:   "11122223333",
				Password: "123456",
				Avatar:   "jisoo.jpeg",
				Sex:      1,
			},
			expect: true,
		},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(e.args)
			if err != nil {
				t.Fatal(err)
			}

			if e.expect {
				t.Log(got.Id)
			}
		})
	}
}