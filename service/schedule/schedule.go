package schedule

import (
	"box/base"
	"box/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type RetrieveInput struct {
	BeginTime int64 `json:"beginTime"`
	EndTime   int64 `json:"endTime"`
}

func Retrieve(ctx *gin.Context, input RetrieveInput) (map[string]interface{}, error) {
	userID := ctx.GetInt64("_userID")
	totalCount, schedules, err := model.ScheduleDao.RetrieveSchedules(ctx, userID, input.BeginTime, input.EndTime)
	if err != nil {
		return nil, base.GetErrorWithMsg("检索日程失败")
	}
	type schedule struct {
		ID        int64  `json:"id"`        // 自增ID
		Title     string `json:"title"`     // 日程标题
		Content   string `json:"content"`   // 日程内容
		BeginTime int64  `json:"beginTime"` // 日程开始时间
		EndTime   int64  `json:"endTime"`   // 日程结束时间
		UpdatedAt int64  `json:"updatedAt"` // 更新时间
	}
	output := make([]schedule, 0, len(schedules))
	for _, v := range schedules {
		output = append(output, schedule{
			ID:        v.ID,
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
	Content   string `json:"content" binding:"required"`
	BeginTime int64  `json:"beginTime" binding:"required"`
	EndTime   int64  `json:"endTime" binding:"required"`
}

func Create(ctx *gin.Context, input CreateInput) error {
	userID := ctx.GetInt64("_userID")
	schedule := model.Schedule{
		UserID:    userID,
		Title:     input.Title,
		Content:   input.Content,
		BeginTime: input.BeginTime,
		EndTime:   input.EndTime,
	}
	err := model.ScheduleDao.CreateRecord(ctx, schedule)
	if err != nil {
		log.WithField("schedule", schedule).Errorf("schedule create failed, err: %s", err.Error())
		return base.GetErrorWithMsg("日程创建失败")
	}
	return nil
}
