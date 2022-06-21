package output

import (
	"box/base"
	validator2 "box/base/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
)

type output struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func Success(ctx *gin.Context, data map[string]interface{}) {
	if data == nil || &data == nil {
		data = map[string]interface{}{}
	}
	ctx.JSON(http.StatusOK, output{
		Code: 0,
		Msg:  "成功",
		Data: data,
	})
}

func Failure(ctx *gin.Context, err error) {
	//log.WithContext(ctx).Errorf("%+v", err)
	code, msg := -1, errors.Cause(err).Error()
	switch errors.Cause(err).(type) {
	case base.Error:
		code = errors.Cause(err).(base.Error).Code
		msg = errors.Cause(err).(base.Error).Msg
	case validator.ValidationErrors:
		code = base.ErrorCodeInvalidParam
		msg = validator2.TranslateValidatorError(err.(validator.ValidationErrors))
	default:
	}
	ctx.JSON(http.StatusOK, output{
		Code: code,
		Msg:  msg,
		Data: map[string]interface{}{},
	})
}
