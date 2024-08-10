package xerr

var message = map[uint32]string{
	OK:                  "请求响应成功",
	SERVER_COMMON_ERROR: "服务器开小差啦，稍后再来试一试",
	REUQEST_PARAM_ERROR: "请求参数错误",
	DB_ERROR:            "数据库繁忙，请稍后再试",
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦，稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
