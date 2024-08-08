package util

import (
	"context"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "identity"

func GetUidFromCtx(ctx context.Context) string {
	uid, ok := ctx.Value(CtxKeyJwtUserId).(string)
	if ok {
		return uid
	} else {
		return ""
	}
}
