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
	Email    string `gorm:"column:email;type:varchar(50);not null;uniqueIndex:uk_email,priority:1" json:"email"` // 电子邮箱
	Password string `gorm:"column:password;type:varchar(20);not null" json:"password"`                           // 密码
	gorm.Model
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

func (d userDao) GetRecords(ctx *gin.Context, options ...func(db *gorm.DB) *gorm.DB) ([]User, error) {
	var records []User
	result := preload.DB.WithContext(ctx).Scopes(options...).Find(&records)
	if result.Error != nil {
		return records, errors.Wrapf(base.ErrorDBSelect, "get record fail, err: %s", result.Error.Error())
	}
	return records, nil
}

func (d userDao) GetRecord(ctx *gin.Context, options ...func(db *gorm.DB) *gorm.DB) (User, error) {
	var record User
	result := preload.DB.WithContext(ctx).Scopes(options...).Find(&record).Limit(1)
	if result.Error != nil {
		return record, errors.Wrapf(base.ErrorDBSelect, "get record fail, err: %s", result.Error.Error())
	}
	return record, nil
}

//func (d userDao) RetrieveUserByEmail(ctx *gin.Context, email string) (*User, error) {
//	var user User
//	result := preload.DB.WithContext(ctx).Where(User{Email: email}).Find(&user).Limit(1)
//	if result.Error != nil {
//		return nil, errors.Wrapf(result.Error, "")
//	}
//	return &user, nil
//}

func (d userDao) AddRecord(ctx *gin.Context, email string, password string) (User, error) {
	user := User{
		Email:    email,
		Password: password,
	}
	result := preload.DB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return User{}, errors.Wrapf(base.ErrorDBInsert, "add record fail, err: %s", result.Error.Error())
	}
	return user, nil
}
