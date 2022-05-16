package model

import (
	"box/base"
	"box/preload"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

func (d medicalDao) GetRecords(ctx *gin.Context, options ...func(db *gorm.DB) *gorm.DB) (int64, []Medical, error) {
	var data []Medical
	db := preload.DB.WithContext(ctx).Scopes(options...)
	var totalCount int64
	result := db.Find(&data).Offset(-1).Limit(-1).Count(&totalCount)
	if result.Error != nil {
		return 0, nil, errors.Wrapf(base.ErrorDBSelect, "get records fail, err: %s", result.Error.Error())
	}
	return totalCount, data, nil
}

func (d medicalDao) AddRecord(ctx *gin.Context, value Medical) error {
	result := preload.DB.WithContext(ctx).Create(&value)
	if result.Error != nil {
		return errors.Wrapf(base.ErrorDBInsert, "add record fail, err: %s", result.Error.Error())
	}
	return nil
}

func (d medicalDao) UpdateRecord(ctx *gin.Context, condition Medical, newValue Medical) error {
	result := preload.DB.WithContext(ctx).Where(condition).Updates(&newValue)
	if result.Error != nil {
		return errors.Wrapf(base.ErrorDBUpdate, "update record fail, err: %s", result.Error.Error())
	}
	return nil
}
