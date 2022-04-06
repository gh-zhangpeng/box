package medical

import (
	"box/base"
	"box/base/output"
	"box/service/ofy/medical"
	"github.com/gin-gonic/gin"
)

func AddRecord(ctx *gin.Context) {
	var input medical.AddRecordInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	err := medical.AddRecord(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, map[string]interface{}{})
}

func GetRecords(ctx *gin.Context) {
	var input medical.GetRecordsInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, base.ErrorInvalidParam)
		return
	}
	data, err := medical.GetRecords(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, data)
}
