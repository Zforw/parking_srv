package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
	"parking/global"
)

func InitValidator() {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("翻译器初始化错误")
		return
	}
}

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New()
		//参数1：备用语言，后续：要支持的语言
		uni := ut.New(enT, zhT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translation.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			zh_translation.RegisterDefaultTranslations(v, global.Trans)
		default:
			en_translation.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}
