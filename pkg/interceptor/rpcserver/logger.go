package rpcserver

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"im/pkg/xerr"
)

// gRPC server 拦截器
func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err == nil {
		return resp, nil
	}
	
	logx.WithContext(ctx).Errorf("【RPC SRV ERR】 %v", err)

	causeErr := errors.Cause(err)
	if e, ok := causeErr.(*xerr.CodeMsg); ok {
		err = status.Error(codes.Code(e.GetCode()), e.GetMsg())  // 转成 gRPC error
	}

	return resp, err
}
