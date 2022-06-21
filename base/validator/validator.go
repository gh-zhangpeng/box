package validator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Init() {
	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	validate = binding.Validator.Engine().(*validator.Validate)
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(fmt.Errorf("init validator failed: %w \n", err))
	}
}

func TranslateValidatorError(errs validator.ValidationErrors) string {
	var errsMsgs []string
	for _, v := range errs.Translate(trans) {
		errsMsgs = append(errsMsgs, v)
	}
	return strings.Join(errsMsgs, ",")
}
