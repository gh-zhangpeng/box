package schedule

import (
	"box/base"
	"box/base/output"
	"box/service/schedule"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	var input struct {
		ID int64 `json:"ID"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	err := schedule.Delete(ctx, input.ID)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, nil)
}

func Update(ctx *gin.Context) {
	var input schedule.UpdateInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	data, err := schedule.Update(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, data)
}

func Retrieve(ctx *gin.Context) {
	var input schedule.RetrieveInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	data, err := schedule.Retrieve(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, data)
}

func Create(ctx *gin.Context) {
	var input schedule.CreateInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	err := schedule.Create(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, nil)
}
