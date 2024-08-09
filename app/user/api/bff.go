package main

import (
	"flag"
	"fmt"

	"im/app/user/api/internal/config"
	"im/app/user/api/internal/handler"
	"im/app/user/api/internal/svc"
	"im/pkg/serializer"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/bff-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册错误处理函数
	httpx.SetErrorHandlerCtx(serializer.ErrHandler(c.Name))
	httpx.SetOkHandler(serializer.OkHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
