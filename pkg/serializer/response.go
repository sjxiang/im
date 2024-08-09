package serializer


// 序列化器

type Response struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func FeedbackOk(data any) *Response {
	return &Response{
		Code: 200,
		Msg:  "OK",
		Data: data,
	}
}

func FeedbackFail(errCode uint32, errMsg string) *Response {
	return &Response{
		Code: errCode,
		Msg:  errMsg,
	}
}
