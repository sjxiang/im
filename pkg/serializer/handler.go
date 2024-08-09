package serializer

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"

	"im/pkg/xerr"
)

// 成功返回
func OkHandler(_ context.Context, v interface{}) any {
	return FeedbackOk(v)
}

func ErrHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
	
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // 提取出原始 err

		// 错误类型 1、当前 api 服务传递的
		if e, ok := causeErr.(*xerr.CodeError); ok { 
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			// 错误类型 2、上游 gRPC 服务传递的
			if grpcStatus, ok := status.FromError(causeErr); ok { 
				
				errcode = uint32(grpcStatus.Code())
				errmsg = grpcStatus.Message()
			}
			
			// 错误类型 3、未知错误
		}

		// 日志记录
		logx.WithContext(ctx).Errorf("【%s】_【%+v】", name, err)

		return http.StatusBadRequest, FeedbackFail(errcode, errmsg)
	}
}