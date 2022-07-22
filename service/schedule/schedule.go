package schedule

import (
	"box/base"
	"box/dal/model"
	"box/dal/query"
	"box/middleware"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gen"
	"time"
)

func Delete(ctx *gin.Context, ID int64) error {
	schedule := query.Schedule
	_, err := schedule.WithContext(ctx).
		Where(schedule.ID.Eq(ID), schedule.DeletedAt.Eq(0)).
		UpdateColumn(schedule.DeletedAt, time.Now().Unix())
	if err != nil {
		log.WithField("ID", ID).Errorf("schedule delete fail, err: %s", err.Error())
		return base.ErrorSystemError
	}
	return nil
}

type UpdateInput struct {
	ID        int64  `json:"ID" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content"`
	BeginTime int64  `json:"beginTime"`
	EndTime   int64  `json:"endTime"`
}

func Update(ctx *gin.Context, input UpdateInput) error {
	newValue := model.Schedule{
		Title:     input.Title,
		Content:   input.Content,
		BeginTime: input.BeginTime,
		EndTime:   input.EndTime,
		UpdatedAt: time.Now().Unix(),
	}
	schedule := query.Schedule
	_, err := schedule.WithContext(ctx).
		Select(schedule.Title, schedule.Content, schedule.BeginTime, schedule.EndTime, schedule.UpdatedAt).
		Where(schedule.ID.Eq(input.ID), schedule.DeletedAt.Eq(0)).
		UpdateColumns(newValue)
	if err != nil {
		log.WithField("ID", input.ID).WithField("newValue", newValue).Errorf("schedule update fail, struct to map fail, err: %s", err.Error())
		return base.ErrorSystemError
	}
	return nil
}

type RetrieveInput struct {
	BeginTime int64 `form:"beginTime"`
	EndTime   int64 `form:"endTime"`
	PageNo    int   `form:"pageNo" binding:"min=1"`
	PageSize  int   `form:"pageSize" binding:"min=1"`
}

func Retrieve(ctx *gin.Context, input RetrieveInput) (map[string]interface{}, error) {
	userID := middleware.GetUserID(ctx)
	schedule := query.Schedule
	conditions := make([]gen.Condition, 0, 5)
	conditions = append(conditions, schedule.UserID.Eq(userID))
	conditions = append(conditions, schedule.DeletedAt.Eq(0))
	if input.BeginTime > 0 {
		conditions = append(conditions, schedule.BeginTime.Gte(input.BeginTime))
	}
	if input.EndTime > 0 {
		conditions = append(conditions, schedule.EndTime.Lte(input.EndTime))
	}
	records, totalCount, err := schedule.WithContext(ctx).
		Where(conditions...).
		FindByPage((input.PageNo-1)*input.PageSize, input.PageSize)
	if err != nil {
		return nil, errors.Wrapf(base.GetErrorWithMsg("检索日程失败"), "err: %s", err.Error())
	}
	type tempSchedule struct {
		ID        int64  `json:"id"`        // 自增ID
		UserID    int64  `json:"userID"`    // 用户ID
		Title     string `json:"title"`     // 日程标题
		Content   string `json:"content"`   // 日程内容
		BeginTime int64  `json:"beginTime"` // 日程开始时间
		EndTime   int64  `json:"endTime"`   // 日程结束时间
		UpdatedAt int64  `json:"updatedAt"` // 更新时间
	}
	output := make([]tempSchedule, 0, len(records))
	for _, v := range records {
		output = append(output, tempSchedule{
			ID:        v.ID,
			UserID:    v.UserID,
			Title:     v.Title,
			Content:   v.Content,
			BeginTime: v.BeginTime,
			EndTime:   v.EndTime,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return map[string]interface{}{
		"totalCount": totalCount,
		"schedules":  output,
	}, nil
}

type CreateInput struct {
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content"`
	BeginTime int64  `json:"beginTime"`
	EndTime   int64  `json:"endTime"`
}

func Create(ctx *gin.Context, input CreateInput) error {
	newValue := model.Schedule{
		UserID:    middleware.GetUserID(ctx),
		Title:     input.Title,
		Content:   input.Content,
		BeginTime: input.BeginTime,
		EndTime:   input.EndTime,
		CreatedAt: time.Now().Unix(),
	}
	schedule := query.Schedule
	err := schedule.WithContext(ctx).Omit(schedule.UpdatedAt).Create(&newValue)
	if err != nil {
		log.WithField("schedule", schedule).Errorf("schedule create fail, err: %s", err.Error())
		return errors.Wrapf(base.GetErrorWithMsg("日程创建失败"), "err: %s", err.Error())
	}
	return nil
}
