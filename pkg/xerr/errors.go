package xerr

import (
	"fmt"
)


/*
	通用错误 "github.com/zeromicro/x/errors"
 */

type CodeMsg struct {
	Code uint32  
	Msg  string
}

func (c *CodeMsg) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", c.Code, c.Msg)
}

// 快速创建错误对象
func NewCodeMsg(code uint32, msg string) error {
	return &CodeMsg{Code: code, Msg: msg}
}

// 返回给前端的错误码
func (e *CodeMsg) GetCode() uint32 {
	return e.Code
}

// 返回给前端显示端错误信息
func (e *CodeMsg) GetMsg() string {
	return e.Msg
}

func NewDBErr() error {
	return &CodeMsg{Code: DB_ERROR, Msg: MapErrMsg(DB_ERROR)}
}

func NewInternalServerErr() error {
	return &CodeMsg{Code: SERVER_COMMON_ERROR, Msg: MapErrMsg(SERVER_COMMON_ERROR)}
}

