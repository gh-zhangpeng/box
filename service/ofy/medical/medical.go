package medical

import (
	"box/base"
	"box/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UpdateInput struct {
	ID                uint    `json:"ID" binding:"required"`
	Height            float32 `json:"height"`
	Weight            float32 `json:"weight"`
	HeadCircumference float32 `json:"headCircumference"`
}

func UpdateRecord(ctx *gin.Context, input UpdateInput) error {
	if input.Height <= 0 && input.Weight <= 0 && input.HeadCircumference <= 0 {
		return base.ErrorInvalidParam
	}
	condition := model.Medical{
		Model: gorm.Model{ID: input.ID},
	}
	newValue := model.Medical{
		Height:            input.Height,
		Weight:            input.Weight,
		HeadCircumference: input.HeadCircumference,
		Model: gorm.Model{
			UpdatedAt: time.Now(),
		},
	}
	err := model.MedicalDao.UpdateRecord(ctx, condition, newValue)
	if err != nil {
		log.WithField("condition", condition).WithField("newValue", newValue).Errorf("medical update record fail, err: %s", err.Error())
		return base.ErrorSystemError
	}
	return nil
}

type AddInput struct {
	Height            float32 `json:"height"`
	Weight            float32 `json:"weight"`
	HeadCircumference float32 `json:"headCircumference"`
}

func AddRecord(ctx *gin.Context, input AddInput) error {
	if input.Height <= 0 && input.Weight <= 0 && input.HeadCircumference <= 0 {
		return base.ErrorInvalidParam
	}
	medical := model.Medical{
		Height:            input.Height,
		Weight:            input.Weight,
		HeadCircumference: input.HeadCircumference,
	}
	err := model.MedicalDao.AddRecord(ctx, medical)
	if err != nil {
		log.WithField("medical", medical).Errorf("medical add record fail, err: %s", err.Error())
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
		log.Errorf("medical get records fail, err: %s", err.Error())
		return nil, err
	}
	type medical struct {
		ID                uint    `json:"ID"`
		Height            float32 `json:"height"`
		Weight            float32 `json:"weight"`
		HeadCircumference float32 `json:"headCircumference"`
		UpdatedAt         string  `json:"updatedAt"`
	}
	medicals := make([]medical, 0, len(data))
	for _, v := range data {
		medicals = append(medicals, medical{
			ID:                v.ID,
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
