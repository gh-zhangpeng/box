package schedule

import (
	"box/base"
	"box/base/output"
	"box/service/schedule"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Retrieve(ctx *gin.Context) {
	var input schedule.RetrieveInput
	if err := ctx.ShouldBind(&input); err != nil {
		log.WithField("input", input).Errorf("schedule retrieve failed, err: %s", err.Error())
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	data, err := schedule.Retrieve(ctx, input)
	if err != nil {
		log.Errorf("schedule retrieve failed, err: %s", err.Error())
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, data)
}

func Create(ctx *gin.Context) {
	var input schedule.CreateInput
	if err := ctx.ShouldBind(&input); err != nil {
		log.WithField("input", input).Errorf("schedule create failed, err: %s", err.Error())
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	err := schedule.Create(ctx, input)
	if err != nil {
		log.WithField("input", input).Errorf("schedule create failed, err: %s", err.Error())
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, nil)
}
