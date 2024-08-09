package xerr


/*

				用户反馈
               /
系统 -> 错误 ->
		       \
			    中间件 -> 日志记录
			
 */


// 成功返回
const OK uint32 = 200

// 全局错误码
const (
	SERVER_COMMON_ERROR uint32 = 100001
	REUQEST_PARAM_ERROR uint32 = 100002 
	DB_ERROR            uint32 = 100003
	DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006

	TOKEN_INVALID_ERROR uint32 = 100004
	TOKEN_EXPIRE_ERROR  uint32 = 100005
	TOKEN_GENERATE_ERROR uint32 = 100005
)

/*** 前3位代表业务，后3位代表具体功能 ***/

// 用户模块
