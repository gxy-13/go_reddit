package controller

import (
	"fmt"

	en2 "github.com/go-playground/validator/v10/translations/en"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// 定义一个全局翻译器
var trans ut.Translator

// InitTrans
func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器

		// 第一个参数是备用的语言环境
		// 后面的参数是应该支持的语言环境，支持多个
		uni := ut.New(enT, zhT, enT)

		// locale通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator()传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = en2.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = en2.RegisterDefaultTranslations(v, trans)
		default:
			err = en2.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}
