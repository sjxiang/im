package util

import (
	"context"
)

var CtxKeyJwtUserId = "identity"

// 从 context.Context 中提取用户 ID
func GetUidFromCtx(ctx context.Context) string {
	// 尝试从上下文中获取 indentity 对应的值，并断言为字符串类型
	uid, ok := ctx.Value(CtxKeyJwtUserId).(string)
	if ok {
		// 如果断言成功，返回获取到的字符串值
		return uid
	} else {
		// 如果断言失败，返回空字符串
		return ""
	}
}