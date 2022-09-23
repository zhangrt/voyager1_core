package validate

import (
	"errors"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	idvalidator "github.com/guanguans/id-validator"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// 定义一个全局翻译器T
var trans ut.Translator

// 出生日期（不晚于当前日期）
func checkBirthDateFunc(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return true
	}
	return false
}

// 身份证校验规则(用于struct)
func checkCardIdFunc(fl validator.FieldLevel) bool {
	//if country == "" || strings.ToUpper(country) == "CN" {
	id := fl.Field().String()
	//idReg := regexp.MustCompile(`^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$|^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}([0-9]|X)$`)
	//return idReg.MatchString(id)
	return idvalidator.IsValid(id, true)
	//} else {
	//	return false
	//}
}

// 电话校验规则(用于struct)
func checkPhoneFunc(fl validator.FieldLevel) bool {
	//if country == "" || strings.ToUpper(country) == "CN" {
	phone := fl.Field().String()
	reg := regexp.MustCompile(`(^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|195|198|199|(147))\\d{8}$)|([0-9]{3,4}[0-9]{8}$)`)
	return reg.MatchString(phone)
	//} else {
	//	return false
	//}
}

// 护照号码校验
func checkPassportNoFunc(fl validator.FieldLevel) bool {
	//if country == "" || strings.ToUpper(country) == "CN" {
	portNo := fl.Field().String()
	reg := regexp.MustCompile(`((^[E,K])[0-9]{8}$)|((^(SE)|^(DE)|^(PE)|^(MA))[0-9]{7}$)`)
	return reg.MatchString(portNo)
	//} else {
	//	return false
	//}
}

/*// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

// InitTrans 初始化翻译器
func InitTrans(locale string, model any) (err error) {
	validate := validator.New()
	// 注册自定义校验
	validate.RegisterValidation("beforeCurrentDate", checkBirthDateFunc)
	validate.RegisterValidation("isCardId", checkCardIdFunc)
	validate.RegisterValidation("isPhoneNo", checkPhoneFunc)
	validate.RegisterValidation("isPassportNo", checkPassportNoFunc)
	zhT := zh.New() // 中文翻译器
	//enT := en.New() // 英文翻译器
	uni := ut.New(zhT, zhT)
	var ok bool
	trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}
	// 注册翻译器
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(validate, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(validate, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(validate, trans)
	}
	if err != nil {
		return err
	}
	// 注意！因为这里会使用到trans实例
	// 所以这一步注册要放到trans初始化的后面
	if err := validate.RegisterTranslation(
		"beforeCurrentDate",
		trans,
		registerTranslator("beforeCurrentDate", "{0}不能晚于当前日期"),
		translate,
	); err != nil {
		return err
	}

	if err := validate.RegisterTranslation(
		"isCardId",
		trans,
		registerTranslator("isCardId", "{0}不正确"),
		translate,
	); err != nil {
		return err
	}

	if err := validate.RegisterTranslation(
		"isPhoneNo",
		trans,
		registerTranslator("isPhoneNo", "{0}不正确"),
		translate,
	); err != nil {
		return err
	}

	if err := validate.RegisterTranslation(
		"isPassportNo",
		trans,
		registerTranslator("isPassportNo", "{0}不正确"),
		translate,
	); err != nil {
		return err
	}
	err = validate.Struct(model)
	if err != nil {
		return err
	}
	return
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}*/

/*const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"
	//locale = "zh"
)*/

// 自定义翻译器
func TransInit(model interface{}, locale string) error {
	//c.ShouldBind(model)
	//var locale = "zh" // todo
	zhT := zh.New()
	enT := en.New()
	uni := ut.New(zhT, zhT, enT)
	validate := validator.New()
	trans, _ := uni.GetTranslator(locale)
	// 注册翻译器
	switch locale {
	case "zh":
		zhTranslations.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("comment")
		})
		validate.RegisterValidation("beforeCurrentDate", checkBirthDateFunc)
		validate.RegisterTranslation("beforeCurrentDate", trans, func(ut ut.Translator) error {
			return ut.Add("beforeCurrentDate", "{0}不能晚于当前日期", true)
		},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("beforeCurrentDate", fe.Field())
				return t
			},
		)

		validate.RegisterValidation("isCardId", checkCardIdFunc)
		validate.RegisterTranslation("isCardId", trans, func(ut ut.Translator) error {
			return ut.Add("isCardId", "{0}不正确", true)
		},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("isCardId", fe.Field())
				return t
			},
		)

		validate.RegisterValidation("isPhoneNo", checkPhoneFunc)
		validate.RegisterTranslation("isPhoneNo", trans, func(ut ut.Translator) error {
			return ut.Add("isPhoneNo", "{0}不正确", true)
		},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("isPhoneNo", fe.Field())
				return t
			},
		)

		validate.RegisterValidation("isPassportNo", checkPassportNoFunc)
		validate.RegisterTranslation("isPassportNo", trans, func(ut ut.Translator) error {
			return ut.Add("isPassportNo", "{0}不正确", true)
		},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("isPassportNo", fe.Field())
				return t
			},
		)

	case "en":
		enTranslations.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("en_comment")
		})
	default:
		zhTranslations.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("comment")
		})
	}
	//c.Set(TranslatorKey, trans)
	//c.Set(ValidatorKey, validate)
	err := validate.Struct(model)
	// 错误信息以切片形式输出
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))

	}
	return nil
}

/*func DefaultGetValid(c *gin.Context, model interface{}) error {
	c.ShouldBind(model)
	// 获取验证器
	validate, _ := c.Get(ValidatorKey)
	valid, _ := validate.(*validator.Validate)
	// 获取翻译器
	tran, _ := c.Get(TranslatorKey)
	trans, _ := tran.(ut.Translator)
	err := valid.Struct(model)
	// 错误信息以切片形式输出
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}*/
