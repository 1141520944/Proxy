package mildware

import (
	"strings"

	"github.com/1141520944/proxy/server/pkg/models"
	"github.com/1141520944/proxy/server/pkg/util/jwt"
	"github.com/1141520944/proxy/server/pkg/util/request"
	"github.com/gin-gonic/gin"
)

var result models.ResultData

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			// c.JSON(http.StatusOK, gin.H{
			// 	"code": 2003,
			// 	"msg":  "请求头中auth为空",
			// })
			result.ResponseError(c, models.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// c.JSON(http.StatusOK, gin.H{
			// 	"code": 2004,
			// 	"msg":  "请求头中auth格式有误",
			// })
			result.ResponseError(c, models.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			// c.JSON(http.StatusOK, gin.H{
			// 	"code": 2005,
			// 	"msg":  "无效的Token",
			// })
			result.ResponseError(c, models.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(request.CtxUserIDkey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get(ctxUserIDkey)来获取当前请求的用户信息
	}
}
