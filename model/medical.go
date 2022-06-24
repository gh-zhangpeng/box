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
	ID                int64   `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`              // 自增主键
	Height            float32 `gorm:"column:height;type:float unsigned;not null" json:"height"`                         // 身高，单位：厘米
	Weight            float32 `gorm:"column:weight;type:float unsigned;not null" json:"weight"`                         // 体重，单位：公斤
	HeadCircumference float32 `gorm:"column:head_circumference;type:float unsigned;not null" json:"head_circumference"` // 头围
	OperatorID        int64   `gorm:"column:operator_id;type:int;not null" json:"operator_id"`                          // 操作人id
	CreatedAt         int64   `gorm:"column:created_at;type:int;not null" json:"created_at"`                            // 创建时间
	UpdatedAt         int64   `gorm:"column:updated_at;type:int;not null" json:"updated_at"`                            // 更新时间
	DeletedAt         int64   `gorm:"column:deleted_at;type:int;not null" json:"deleted_at"`                            // 删除时间
}

// TableName Medical's table name
func (*Medical) TableName() string {
	return tableNameMedical
}

var MedicalDao medicalDao

type medicalDao struct{}

func (d medicalDao) CreateRecord(ctx *gin.Context, value Medical) error {
	result := preload.DB.WithContext(ctx).Create(&value)
	if result.Error != nil {
		return errors.Wrapf(base.ErrorDBInsert, "add record failed, err: %s", result.Error.Error())
	}
	return nil
}

func (d medicalDao) UpdateRecordByID(ctx *gin.Context, ID int64, newValue Medical) error {
	result := preload.DB.WithContext(ctx).Model(&Medical{}).Where(Medical{ID: ID, DeletedAt: 0}).Limit(1).Updates(&newValue)
	if result.Error != nil {
		return errors.Wrapf(base.ErrorDBUpdate, "update record failed, err: %s", result.Error.Error())
	}
	return nil
}

func (d medicalDao) UpdateRecordByIDWithMap(ctx *gin.Context, ID int64, newValue map[string]interface{}) error {
	result := preload.DB.WithContext(ctx).Model(&Medical{}).Where(&Medical{ID: ID, DeletedAt: 0}).Limit(1).Updates(newValue)
	if result.Error != nil {
		return errors.Wrapf(base.ErrorDBUpdate, "update record failed, err: %s", result.Error.Error())
	}
	return nil
}

func (d medicalDao) RetrieveRecords(ctx *gin.Context, options ...func(db *gorm.DB) *gorm.DB) (int64, []Medical, error) {
	var data []Medical
	db := preload.DB.WithContext(ctx).Scopes(options...)
	var totalCount int64
	result := db.Find(&data).Offset(-1).Limit(-1).Count(&totalCount)
	if result.Error != nil {
		return 0, nil, errors.Wrapf(base.ErrorDBSelect, "retrieve records failed, err: %s", result.Error.Error())
	}
	return totalCount, data, nil
}
