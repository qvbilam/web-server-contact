package middleware

import (
	"contact/api"
	userProto "contact/api/qvbilam/user/v1"
	"contact/global"
	"context"
	"github.com/gin-gonic/gin"
)

// Auth 验证jwt
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := global.UserServerClient.Auth(context.Background(), &userProto.AuthRequest{
			Token: ctx.Request.Header.Get("Authorization"),
		})

		if err != nil {
			api.HandleGrpcErrorToHttp(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Set("userId", user.Id)
		// 继续执行
		ctx.Next()
	}
}
