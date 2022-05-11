package model

import (
	"box/preload"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

const tableNameSchedule = "tblSchedule"

// Schedule mapped from table <tblSchedule>
type Schedule struct {
	UserID    int32     `gorm:"column:user_id;type:int unsigned;not null" json:"user_id"`      // 用户ID
	Title     string    `gorm:"column:title;type:varchar(255);not null" json:"title"`          // 日程标题
	Content   string    `gorm:"column:content;type:varchar(500)" json:"content"`               // 日程内容
	BeginTime time.Time `gorm:"column:begin_time;type:datetime(3);not null" json:"begin_time"` // 日程开始时间
	EndTime   time.Time `gorm:"column:end_time;type:datetime(3);not null" json:"end_time"`     // 日程结束时间
	gorm.Model
}

// TableName Schedule's table name
func (*Schedule) TableName() string {
	return tableNameSchedule
}

var ScheduleDao scheduleDao

type scheduleDao struct{}

func (d scheduleDao) RetrieveSchedules(ctx *gin.Context, userID int32, beginTime, endTime int64) ([]Schedule, error) {
	var schedules []Schedule
	query := preload.DB.WithContext(ctx).Where(Schedule{UserID: userID})
	if beginTime > 0 {
		query = query.Where("begin_time >= ?", time.UnixMilli(beginTime))
	}
	if endTime > 0 {
		query = query.Where("end_time <= ?", time.UnixMilli(endTime))
	}
	result := query.Find(&schedules)
	if result.Error != nil {
		log.Errorf("retrieve schedules fail, err: %s, userID: %d, beginTime: %+d, endTime: %d", result.Error.Error(), userID, beginTime, endTime)
		return nil, result.Error
	}
	return schedules, nil
}

func (d scheduleDao) CreateSchedule(ctx *gin.Context, schedule Schedule) (*Schedule, error) {
	result := preload.DB.WithContext(ctx).Create(&schedule)
	if result.Error != nil {
		log.Errorf("create schedule fail, err: %s, schedule: %+v", result.Error.Error(), schedule)
		return nil, result.Error
	}
	return &schedule, nil
}
