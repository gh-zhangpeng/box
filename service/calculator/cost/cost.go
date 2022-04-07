package cost

import (
	"github.com/gin-gonic/gin"
)

type SubwayCostInput struct {
	Times     int `form:"times"`
	UnitPrice int `form:"unitPrice" binding:"required"`
}

func GetSubwayCost(ctx *gin.Context, param SubwayCostInput) (float32, error) {
	total := float32(param.Times * param.UnitPrice)
	if total <= 100 {
		return total, nil
	}
	const (
		discountLadderFirst  = 100
		discountLadderSecond = 150
		discountLadderThird  = 400

		discountRateDefault = 1
		discountRateFirst   = 0.8
		discountRateSecond  = 0.5
	)
	total = 0
	remainTimes := 0
	for i := 0; i < param.Times; i++ {
		total += float32(param.UnitPrice) * discountRateDefault
		if total > discountLadderFirst {
			remainTimes = param.Times - i
			break
		}
	}
	for i := 0; i < remainTimes; i++ {
		total += float32(param.UnitPrice) * discountRateFirst
		if total > discountLadderSecond {
			remainTimes = remainTimes - i
			break
		}
	}
	for i := 0; i < remainTimes; i++ {
		total += float32(param.UnitPrice) * discountRateSecond
		if total > discountLadderThird {
			remainTimes = remainTimes - i
			break
		}
	}
	total = total + float32(remainTimes*param.UnitPrice*discountRateDefault)
	return total, nil
}
