package schedule

import (
	"github.com/gin-gonic/gin"
)

type RetrieveInput struct {
	BeginTime int64 `json:"beginTime"`
	EndTime   int64 `json:"endTime"`
}

func Retrieve(ctx *gin.Context, input RetrieveInput) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
	//userID := ctx.GetUint("_userID")
	//schedules, err := model.ScheduleDao.RetrieveSchedules(ctx, userID, input.BeginTime, input.EndTime)
	//if err != nil {
	//	log.Errorf("retrieve schedules fail, err: %s, userID: %d, beginTime: %+d, endTime: %d", err.Error(), userID, input.BeginTime, input.EndTime)
	//	return nil, base.GetErrorWithMsg("检索日程失败")
	//}
	//type schedule struct {
	//	ID        uint   `json:"id"`        // 自增ID
	//	Title     string `json:"title"`     // 日程标题
	//	Content   string `json:"content"`   // 日程内容
	//	BeginTime int64  `json:"beginTime"` // 日程开始时间
	//	EndTime   int64  `json:"endTime"`   // 日程结束时间
	//	CreatedAt int64  `json:"createdAt"` // 创建时间
	//	//UpdatedAt time.Time      `json:"updatedAt"`               // 更新时间
	//}
	//output := make([]schedule, 0, len(schedules))
	//for _, v := range schedules {
	//	output = append(output, schedule{
	//		ID:        v.ID,
	//		Title:     v.Title,
	//		Content:   v.Content,
	//		BeginTime: v.BeginTime.Unix(),
	//		EndTime:   v.EndTime.Unix(),
	//		CreatedAt: v.CreatedAt.Unix(),
	//	})
	//}
	//return map[string]interface{}{
	//	"schedules": output,
	//}, nil
}

type CreateInput struct {
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	BeginTime int64  `json:"beginTime" binding:"required"`
	EndTime   int64  `json:"endTime" binding:"required"`
}

func Create(ctx *gin.Context, input CreateInput) error {
	return nil
	//userID := ctx.GetUint("_userID")
	//schedule := model.Schedule{
	//	UserID:    userID,
	//	Title:     input.Title,
	//	Content:   input.Content,
	//	BeginTime: time.Unix(input.BeginTime, 0),
	//	EndTime:   time.Unix(input.EndTime, 0),
	//}
	//_, err := model.ScheduleDao.CreateSchedule(ctx, schedule)
	//if err != nil {
	//	log.Errorf("create schedule fail, err: %s, schedule: %+v", err.Error(), schedule)
	//	return base.GetErrorWithMsg("日程创建失败")
	//}
	//return nil
}
