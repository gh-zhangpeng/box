package medical

import (
	"box/base"
	"box/base/output"
	"box/service/medical"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var input medical.CreateInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, err)
		return
	}
	err := medical.Create(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, map[string]interface{}{})
}

func Update(ctx *gin.Context) {
	var input medical.UpdateInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, err)
		return
	}
	err := medical.Update(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, map[string]interface{}{})
}

func Retrieve(ctx *gin.Context) {
	var input medical.RetrieveInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	data, err := medical.Retrieve(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, data)
}

func Delete(ctx *gin.Context) {
	var input medical.RetrieveInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	data, err := medical.Retrieve(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, data)
}
