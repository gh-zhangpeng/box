package account

import (
	"box/base"
	"box/base/jwt"
	"box/model"
	box_lib "github.com/gh-zhangpeng/box-lib"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

func Login(ctx *gin.Context, email, password string) (map[string]interface{}, error) {
	output := map[string]interface{}{}
	user, err := model.UserDao.GetRecord(ctx, model.UserDao.Email(email))
	if err != nil {
		return output, errors.Wrapf(err, "get user fail, email: %s", email)
	}
	if (user == model.User{}) {
		return output, errors.Wrapf(base.GetErrorWithMsg("登陆失败，账户不存在"), "user does not exist, email: %s", email)
	} else if password != user.Password {
		return output, errors.WithStack(base.GetErrorWithMsg("登陆失败，密码错误"))
	}
	expiresAt := time.Now().Add(time.Hour * 12).UnixMilli()
	token, err := jwt.GenerateToken(user.ID, expiresAt)
	if err != nil {
		return output, errors.Wrapf(base.ErrorGenerateToken, "generate token fail, userID: %d", user.ID)
	}
	output["token"] = token
	output["expiresAt"] = expiresAt
	return output, nil
}

func Register(ctx *gin.Context, email, password string) error {
	emailPattern := "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	match, err := box_lib.Match(emailPattern, email)
	if err != nil {
		return errors.Wrapf(base.GetErrorWithMsg("邮件地址格式有误"), "match email fail, email: %s, emailPattern: %s", email, emailPattern)
	}
	if !match {
		return errors.Wrapf(base.GetErrorWithMsg("邮件地址格式有误"), "email format is incorrect, email: %s", email)
	}
	passwordPattern := "^[A-Za-z0-9.]{6,30}$"
	match, err = box_lib.Match(passwordPattern, password)
	if err != nil {
		return errors.Wrapf(base.GetErrorWithMsg("密码格式有误"), "match password fail, password: %s, passwordPattern: %s", password, passwordPattern)
	}
	if !match {
		return errors.Wrapf(base.GetErrorWithMsg("密码格式有误"), "password format is incorrect")
	}
	records, err := model.UserDao.GetRecords(ctx, model.UserDao.Email(email))
	if err != nil {
		return errors.Wrapf(base.GetErrorWithMsg("账户注册失败"), "get user fail, email: %s", email)
	} else if len(records) > 0 {
		return errors.Wrapf(base.GetErrorWithMsg("邮箱已被使用"), "email already exists, email: %s", email)
	}
	_, err = model.UserDao.AddRecord(ctx, email, password)
	if err != nil {
		return errors.Wrap(base.GetErrorWithMsg("账户注册失败"), "add user fail")
	}
	return nil
}
