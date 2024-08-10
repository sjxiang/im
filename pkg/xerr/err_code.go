package xerr


/*

				用户反馈
               /
系统 -> 错误 ->
		       \
			    中间件 -> 日志记录
			
 */




// 全局错误码（infra）
const (
	OK                  uint32 = 200
	SERVER_COMMON_ERROR uint32 = 100001
	REUQEST_PARAM_ERROR uint32 = 100002 
	DB_ERROR            uint32 = 100003
	TOKEN_INVALID_ERROR  uint32 = 200001
	TOKEN_EXPIRE_ERROR   uint32 = 200002
	TOKEN_GENERATE_Failed_ERROR uint32 = 200003
)


/*** 前3位代表业务，后3位代表具体功能 ***/

// 用户模块
const (
	
)
