package account

import (
	"box/base"
	"box/base/output"
	"box/service/ofy/account"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, errors.WithStack(base.ErrorInvalidParam))
		return
	}
	data, err := account.Login(ctx, input.Email, input.Password)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, data)
}

func Register(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		log.WithField("input", input).Errorf("account register fail, err: %s", err.Error())
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	err := account.Register(ctx, input.Email, input.Password)
	if err != nil {
		log.WithField("input", input).Errorf("account register fail, err: %s", err.Error())
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, map[string]interface{}{})
}
