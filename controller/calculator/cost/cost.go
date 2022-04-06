package cost

import (
	"box/base"
	"box/base/output"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCost(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		log.WithField("input", input).Errorf("account register fail, err: %s", err.Error())
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	//err := account.Register(ctx, input.Email, input.Password)
	//if err != nil {
	//	log.WithField("input", input).Errorf("account register fail, err: %s", err.Error())
	//	output.Failure(ctx, err)
	//	return
	//}
	output.Success(ctx, map[string]interface{}{})
}
