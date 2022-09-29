package validate

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	idvalidator "github.com/guanguans/id-validator"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// 出生日期（不晚于当前日期）
func beforeCurrentDate(fl validator.FieldLevel) bool {
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
func isCardId(fl validator.FieldLevel) bool {
	id := fl.Field().String()
	return idvalidator.IsValid(id, true)
}

// 电话校验规则(用于struct)
func isPhoneNo(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	reg := regexp.MustCompile(`(^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|195|198|199|(147))\\d{8}$)|([0-9]{3,4}[0-9]{8}$)`)
	return reg.MatchString(phone)
}

// 护照号码校验
func isPassportNo(fl validator.FieldLevel) bool {
	portNo := fl.Field().String()
	reg := regexp.MustCompile(`((^[E,K])[0-9]{8}$)|((^(SE)|^(DE)|^(PE)|^(MA))[0-9]{7}$)`)
	return reg.MatchString(portNo)
}

// 自定义翻译器
func TransInit(model interface{}, locale string) error {
	zhT := zh.New()
	enT := en.New()
	uni := ut.New(zhT, zhT, enT)
	validate := validator.New()
	trans, _ := uni.GetTranslator(locale)
	// 注册自定义校验规则
	validate.RegisterValidation("beforeCurrentDate", beforeCurrentDate)
	validate.RegisterValidation("isCardId", isCardId)
	validate.RegisterValidation("isPhoneNo", isPhoneNo)
	validate.RegisterValidation("isPassportNo", isPassportNo)
	// 注册翻译器
	switch locale {
	case "zh":
		zhTranslations.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("comment")
		})
		ParseConfig("validate_zh.json")
	case "en":
		enTranslations.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("en_comment")
		})
		ParseConfig("validate_en.json")
	default:
		zhTranslations.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("comment")
		})
	}
	RegisterDate(validate, trans, GetConfig().BeforeCurrentDate)
	RegisterCardId(validate, trans, GetConfig().ErrorInfo)
	RegisterPhoneNo(validate, trans, GetConfig().ErrorInfo)
	RegisterPassportNo(validate, trans, GetConfig().ErrorInfo)

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

func RegisterDate(validate *validator.Validate, trans ut.Translator, str string) {
	validate.RegisterTranslation("beforeCurrentDate", trans, func(ut ut.Translator) error {
		return ut.Add("beforeCurrentDate", "{0}"+str, true)
	},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("beforeCurrentDate", fe.Field())
			return t
		},
	)
}

func RegisterCardId(validate *validator.Validate, trans ut.Translator, str string) {
	validate.RegisterTranslation("isCardId", trans, func(ut ut.Translator) error {
		return ut.Add("isCardId", "{0}"+str, true)
	},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("isCardId", fe.Field())
			return t
		},
	)
}

func RegisterPhoneNo(validate *validator.Validate, trans ut.Translator, str string) {
	validate.RegisterTranslation("isPhoneNo", trans, func(ut ut.Translator) error {
		return ut.Add("isPhoneNo", "{0}"+str, true)
	},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("isPhoneNo", fe.Field())
			return t
		},
	)
}

func RegisterPassportNo(validate *validator.Validate, trans ut.Translator, str string) {
	validate.RegisterTranslation("isPassportNo", trans, func(ut ut.Translator) error {
		return ut.Add("isPassportNo", "{0}"+str, true)
	},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("isPassportNo", fe.Field())
			return t
		},
	)
}

var cfg *Config

// 读取配置文件
func ParseConfig(filename string) (*Config, error) {
	expath, _ := os.Getwd()
	file, err := os.Open(expath + "/util/validate/" + filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}

// 定义中间件
func ValidateMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//model:=c.Request.Form
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(zhT, zhT, enT)
		//locale,_ := c.Get("lang") //todo
		locale := "zh"
		trans, _ := uni.GetTranslator(locale)
		validate := validator.New()
		// 注册自定义校验规则
		validate.RegisterValidation("beforeCurrentDate", beforeCurrentDate)
		validate.RegisterValidation("isCardId", isCardId)
		validate.RegisterValidation("isPhoneNo", isPhoneNo)
		validate.RegisterValidation("isPassportNo", isPassportNo)
		// 注册翻译器
		switch locale {
		case "zh":
			zhTranslations.RegisterDefaultTranslations(validate, trans)
			validate.RegisterTagNameFunc(func(field reflect.StructField) string {
				return field.Tag.Get("comment")
			})
			ParseConfig("validate_zh.json")
		case "en":
			enTranslations.RegisterDefaultTranslations(validate, trans)
			validate.RegisterTagNameFunc(func(field reflect.StructField) string {
				return field.Tag.Get("en_comment")
			})
			ParseConfig("validate_en.json")
		default:
			zhTranslations.RegisterDefaultTranslations(validate, trans)
			validate.RegisterTagNameFunc(func(field reflect.StructField) string {
				return field.Tag.Get("comment")
			})
		}
		RegisterDate(validate, trans, GetConfig().BeforeCurrentDate)
		RegisterCardId(validate, trans, GetConfig().ErrorInfo)
		RegisterPhoneNo(validate, trans, GetConfig().ErrorInfo)
		RegisterPassportNo(validate, trans, GetConfig().ErrorInfo)
		//validate.Struct(model)
		c.Set("validateRegister", validate)
		c.Set("translate", trans)
	}
}
