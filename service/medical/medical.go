package medical

import (
	"box/base"
	"box/model"
	box_lib "github.com/gh-zhangpeng/box-lib"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

type UpdateInput struct {
	ID                int64   `json:"ID" binding:"required"`
	Height            float32 `json:"height"`
	Weight            float32 `json:"weight"`
	HeadCircumference float32 `json:"headCircumference"`
}

func Update(ctx *gin.Context, input UpdateInput) error {
	newValue := map[string]interface{}{
		"height":             input.Height,
		"weight":             input.Weight,
		"head_circumference": input.HeadCircumference,
		"updated_at":         time.Now().Unix(),
	}
	err := model.MedicalDao.UpdateRecordByIDWithMap(ctx, input.ID, newValue)
	if err != nil {
		log.WithField("ID", input.ID).WithField("newValue", newValue).Errorf("medical update failed, err: %s", err.Error())
		return base.ErrorSystemError
	}
	return nil
}

type CreateInput struct {
	Height            float32 `json:"height" binding:"required_without_all=Weight HeadCircumference"`
	Weight            float32 `json:"weight" binding:"required_without_all=Height HeadCircumference"`
	HeadCircumference float32 `json:"headCircumference" binding:"required_without_all=Height Weight"`
}

func Create(ctx *gin.Context, input CreateInput) error {
	medical := model.Medical{
		Height:            input.Height,
		Weight:            input.Weight,
		HeadCircumference: input.HeadCircumference,
		OperatorID:        ctx.GetInt64("_userID"),
	}
	err := model.MedicalDao.CreateRecord(ctx, medical)
	if err != nil {
		log.WithField("medical", medical).Errorf("medical add failed, err: %s", err.Error())
		return base.ErrorSystemError
	}
	return nil
}

type RetrieveInput struct {
	PageNo   int `form:"pageNo" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

func Retrieve(ctx *gin.Context, input RetrieveInput) (map[string]interface{}, error) {
	totalCount, records, err := model.MedicalDao.RetrieveRecords(
		ctx, model.Paginate(input.PageNo, input.PageSize),
		model.OrderBy("updated_at desc"),
		model.Deleted(false),
	)
	if err != nil {
		log.Errorf("medical retrieve records failed, err: %s", err.Error())
		return nil, err
	}
	type medical struct {
		ID                int64   `json:"ID"`
		Height            float32 `json:"height"`
		Weight            float32 `json:"weight"`
		HeadCircumference float32 `json:"headCircumference"`
		UpdatedAt         string  `json:"updatedAt"`
		operatorID        int64
		Operator          string `json:"operator"`
	}
	medicals := make([]medical, 0, len(records))
	operatorIDs := make([]uint, 0, len(records))
	for _, record := range records {
		medicals = append(medicals, medical{
			ID:                record.ID,
			Height:            record.Height,
			Weight:            record.Weight,
			HeadCircumference: record.HeadCircumference,
			UpdatedAt:         time.Unix(record.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
			operatorID:        record.OperatorID,
			Operator:          "未知操作人",
		})
		operatorIDs = append(operatorIDs, uint(record.OperatorID))
	}

	users, err := model.UserDao.RetrieveRecords(ctx, model.IDIn(box_lib.UniqueUIntSlice(operatorIDs)))
	if err != nil {
		return nil, err
	}
	userID2Email := make(map[int64]string)
	for _, user := range users {
		userID2Email[user.ID] = user.Email
	}
	for i, v := range medicals {
		if email, ok := userID2Email[v.operatorID]; ok {
			medicals[i].Operator = email
		}
	}
	return map[string]interface{}{
		"totalCount": totalCount,
		"medicals":   medicals,
	}, nil
}
