package util

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/zhangrt/voyager1_core/util/validate"
	"net/http"
)

func test1(c *gin.Context) {
	/*	req := &validate.TestModel{
		CreateDate: time.Date(2022, time.July, 20, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2019, time.July, 20, 0, 0, 0, 0, time.UTC),
		CardId:     "100185201102304434",
		Phone:      "131234567890",
		PassportNo: "DE123456789",
	}*/
	req := &validate.TestModel{}
	c.ShouldBind(req)
	val, _ := c.Get("validateRegister")
	err := val.(*validator.Validate).Struct(req)
	if err != nil {
		trans, _ := c.Get("translate")
		errs := err.(validator.ValidationErrors)
		returnErrs := []validate.ReturnModel{}
		for _, e := range errs {
			returnModel := validate.ReturnModel{Field: e.StructField(), Message: e.Translate(trans.(ut.Translator))}
			returnErrs = append(returnErrs, returnModel)
		}
		c.JSON(404, gin.H{
			"msg": returnErrs,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"model": req,
	})
}

func main() {
	route := gin.Default()
	// 注册中间件
	route.Use(validate.ValidateMiddleWare())
	route.POST("/test", test1)
	route.Run(":8099")
}
