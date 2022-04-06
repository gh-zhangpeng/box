package middleware

import (
	"box/base"
	"box/base/jwt"
	"box/base/output"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil {
			log.Errorf("get token from cookie fail, err: %s", err.Error())
			output.Failure(ctx, base.ErrorSystemError)
			ctx.Abort()
			return
		}
		if len(token) == 0 {
			output.Failure(ctx, base.ErrorNotLogin)
			ctx.Abort()
			return
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				log.Errorf("parse token fail, err: %s", err.Error())
				output.Failure(ctx, err)
				ctx.Abort()
				return
			}
			ctx.Set("_userID", claims.UserID)
		}
		ctx.Next()
	}
}
