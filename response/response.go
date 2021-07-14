package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//封装http返回

//统一格式：
//{
//	code:20001
//	data:xxx
//	msg:xxx
//}

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

//返回成功
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

//返回失败
func Fail(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
