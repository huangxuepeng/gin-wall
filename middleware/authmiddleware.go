package middleware

import (
	"gin-wall/dao"
	"gin-wall/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization") //获取authorization header
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不够"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不够"})
			ctx.Abort()
			return
		}

		userId := claims.UserId //验证通过之后获取claim中的userId
		var user models.UserRegister
		dao.DB.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不够呦",
			})
		}

		ctx.Set("user", user) //用户存在, 将user信息写入Context
		ctx.Next()
	}
}
