package model

import (
	"box/preload"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const tableNameMedical = "tblMedical"

// Medical mapped from table <tblMedical>
type Medical struct {
	Height            float32 `gorm:"column:height;type:float unsigned;default:0" json:"height"`                         // 身高，单位：厘米
	Weight            float32 `gorm:"column:weight;type:float unsigned;default:0" json:"weight"`                         // 体重，单位：公斤
	HeadCircumference float32 `gorm:"column:head_circumference;type:float unsigned;default:0" json:"head_circumference"` // 头围
	gorm.Model
}

// TableName Medical's table name
func (*Medical) TableName() string {
	return tableNameMedical
}

var MedicalDao medicalDao

type medicalDao struct{}

//func (d medicalDao) GetRecord(ctx *gin.Context, options ...func(db *gorm.DB) *gorm.DB) (*model.Project, error) {
//	var data model.Project
//	db := helpers.MISClient.WithContext(ctx).Scopes(options...)
//	result := db.Limit(1).Find(&data)
//	if result.Error != nil {
//		zlog.Errorf(ctx, "GetRecord fail, err: %s", result.Error.Error())
//		return nil, result.Error
//	}
//	return &data, nil
//}

func (d medicalDao) GetRecords(ctx *gin.Context, options ...func(db *gorm.DB) *gorm.DB) (int64, []Medical, error) {
	var data []Medical
	db := preload.DB.WithContext(ctx).Scopes(options...)
	var totalCount int64
	result := db.Find(&data).Offset(-1).Limit(-1).Count(&totalCount)
	if result.Error != nil {
		log.Errorf("medical get records fail, err: %s", result.Error.Error())
		return 0, nil, result.Error
	}
	return totalCount, data, nil
}

func (d medicalDao) AddRecord(ctx *gin.Context, newValue Medical) error {
	result := preload.DB.WithContext(ctx).Create(&newValue)
	if result.Error != nil {
		log.WithField("newValue", newValue).Errorf("medical add record fail, err: %s", result.Error.Error())
		return result.Error
	}
	return nil
}
