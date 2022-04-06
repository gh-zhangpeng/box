package medical

import (
	"box/base"
	"box/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AddRecordInput struct {
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`
}

func AddRecord(ctx *gin.Context, input AddRecordInput) error {
	if input.Height == 0 && input.Weight == 0 {
		return base.ErrorInvalidParam
	}
	fmt.Printf("height: %f, weight: %f\n", input.Height, input.Weight)
	err := model.MedicalDao.AddRecord(ctx, model.Medical{
		Height: input.Height,
		Weight: input.Weight,
	})
	if err != nil {
		return base.ErrorSystemError
	}
	return nil
}

type GetRecordsInput struct {
	PageNo   int `form:"pageNo" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

func GetRecords(ctx *gin.Context, input GetRecordsInput) (map[string]interface{}, error) {
	totalCount, data, err := model.MedicalDao.GetRecords(ctx, model.Paginate(input.PageNo, input.PageSize))
	if err != nil {
		return nil, err
	}
	type medical struct {
		Height            float32 `json:"height"`
		Weight            float32 `json:"weight"`
		HeadCircumference float32 `json:"headCircumference"`
		UpdatedAt         string  `json:"updatedAt"`
	}
	medicals := make([]medical, 0, len(data))
	for _, v := range data {
		medicals = append(medicals, medical{
			Height:            v.Height,
			Weight:            v.Weight,
			HeadCircumference: v.HeadCircumference,
			UpdatedAt:         v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return map[string]interface{}{
		"totalCount": totalCount,
		"medicals":   medicals,
	}, nil
}
