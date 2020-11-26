package middleware

import (
	"fmt"
	"gin_vue/common"
	"gin_vue/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取 authorization header
		tokenStr := ctx.GetHeader("Authorization")
		fmt.Println("请求token ", tokenStr)

		//校验token格式
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token格式不正确",
			})
			ctx.Abort()
			return
		}

		//截取字符
		tokenStr = tokenStr[7:]
		token, claims, err := common.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token格式不正确或者token无效",
			})
			ctx.Abort()
			return
		}

		//	token通过验证，获取claims中的UserID
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		//	验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "用户名或密码不正确",
			})
			ctx.Abort()
			return

		}

		//	用户存在 将用户信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
