package cost

import (
	"box/base"
	"box/base/output"
	"box/service/calculator/cost"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func SubwayCost(ctx *gin.Context) {
	var input cost.SubwayCostInput
	if err := ctx.ShouldBind(&input); err != nil {
		output.Failure(ctx, errors.WithStack(base.ErrorInvalidParam))
		return
	}
	subwayCost, err := cost.GetSubwayCost(ctx, input)
	if err != nil {
		output.Failure(ctx, err)
		return
	}
	output.Success(ctx, map[string]interface{}{
		"cost": subwayCost,
	})
}
