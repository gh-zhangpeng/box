package model

import (
	"box/base"
	"box/preload"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

const tableNameSchedule = "tblSchedule"

// Schedule mapped from table <tblSchedule>
type Schedule struct {
	ID        int64  `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"` // 自增ID
	UserID    int64  `gorm:"column:user_id;type:int unsigned;not null" json:"user_id"`            // 用户ID
	Title     string `gorm:"column:title;type:varchar(255);not null" json:"title"`                // 标题
	Content   string `gorm:"column:content;type:varchar(255);not null" json:"content"`            // 内容
	BeginTime int64  `gorm:"column:begin_time;type:int;not null" json:"begin_time"`               // 开始时间
	EndTime   int64  `gorm:"column:end_time;type:int;not null" json:"end_time"`                   // 结束时间
	CreatedAt int64  `gorm:"column:created_at;type:int;not null" json:"created_at"`               // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int;not null" json:"updated_at"`               // 更新时间
	DeletedAt int64  `gorm:"column:deleted_at;type:int;not null" json:"deleted_at"`               // 删除时间
}

// TableName Schedule's table name
func (*Schedule) TableName() string {
	return tableNameSchedule
}

var ScheduleDao scheduleDao

type scheduleDao struct{}

func (d scheduleDao) RetrieveSchedules(ctx *gin.Context, userID int64, beginTime, endTime int64) (int64, []Schedule, error) {
	var schedules []Schedule
	query := preload.DB.WithContext(ctx).Where(Schedule{UserID: userID})
	if beginTime > 0 {
		query = query.Where("begin_time >= ?", time.UnixMilli(beginTime))
	}
	if endTime > 0 {
		query = query.Where("end_time <= ?", time.UnixMilli(endTime))
	}
	var totalCount int64
	result := query.Find(&schedules).Offset(-1).Limit(-1).Count(&totalCount)
	if result.Error != nil {
		log.Errorf("retrieve schedules failed, err: %s, userID: %d, beginTime: %+d, endTime: %d", result.Error.Error(), userID, beginTime, endTime)
		return 0, nil, errors.Wrapf(base.ErrorDBSelect, "retrieve records failed, err: %s", result.Error.Error())
	}
	return totalCount, schedules, nil
}

func (d scheduleDao) CreateRecord(ctx *gin.Context, schedule Schedule) error {
	result := preload.DB.WithContext(ctx).Create(&schedule)
	if result.Error != nil {
		return errors.Wrapf(base.ErrorDBInsert, "add record failed, err: %s", result.Error.Error())
	}
	return nil
}
