package model

import (
	"box/base"
	"box/preload"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const tableNameUser = "tblUser"

// User mapped from table <tblUser>
type User struct {
	ID        int64  `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`                 // 自增主键
	Email     string `gorm:"column:email;type:varchar(50);not null;uniqueIndex:uk_email,priority:1" json:"email"` // 电子邮箱
	Password  string `gorm:"column:password;type:varchar(20);not null" json:"password"`                           // 密码
	CreatedAt int64  `gorm:"column:created_at;type:int;not null" json:"created_at"`                               // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int;not null" json:"updated_at"`                               // 更新时间
	DeletedAt int64  `gorm:"column:deleted_at;type:int;not null" json:"deleted_at"`                               // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return tableNameUser
}

func (d userDao) Email(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Where("email = ?", email)
		return db
	}
}

var UserDao userDao

type userDao struct{}

func (d userDao) CreateRecord(ctx *gin.Context, email string, password string) (User, error) {
	user := User{
		Email:    email,
		Password: password,
	}
	result := preload.DB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return User{}, errors.Wrapf(base.ErrorDBInsert, "add record failed, err: %s", result.Error.Error())
	}
	return user, nil
}

func (d userDao) RetrieveRecords(ctx *gin.Context, options ...func(db *gorm.DB) *gorm.DB) ([]User, error) {
	var records []User
	result := preload.DB.WithContext(ctx).Scopes(options...).Find(&records)
	if result.Error != nil {
		return records, errors.Wrapf(base.ErrorDBSelect, "retrieve records failed, err: %s", result.Error.Error())
	}
	return records, nil
}

func (d userDao) RetrieveRecord(ctx *gin.Context, options ...func(db *gorm.DB) *gorm.DB) (User, error) {
	var record User
	result := preload.DB.WithContext(ctx).Scopes(options...).Find(&record).Limit(1)
	if result.Error != nil {
		return record, errors.Wrapf(base.ErrorDBSelect, "retrieve record failed, err: %s", result.Error.Error())
	}
	return record, nil
}
