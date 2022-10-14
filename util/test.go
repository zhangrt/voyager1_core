package util

import (
	"github.com/gin-gonic/gin"
	sui18n "github.com/suisrc/gin-i18n"
	"github.com/zhangrt/voyager1_core/util/validate"
	"golang.org/x/text/language"
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
	returnValues := validate.GetReturnValue(c, req)
	if returnValues != nil {
		c.JSON(404, gin.H{
			"msg": returnValues,
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
	/*	// 测试国际化方法
		date := validate.I18nInit("zh", "before_current_date")
		fmt.Println("---------", date)*/
	//在总路由中多语言初始化---国际化中间件
	bundle := sui18n.NewBundle(
		language.Chinese, //默认中文
		"i18n/zh.toml",
		"i18n/en.toml",
	)
	route.Use(sui18n.Serve(bundle))
	//多语言end,添加在具体的路由实例化之前

	// 注册中间件
	route.Use(validate.ValidateMiddleWare())
	route.POST("/test", test1)
	route.Run(":8099")
}
