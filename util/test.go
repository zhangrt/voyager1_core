package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangrt/voyager1_core/util/validate"
	"net/http"
)

func test1(c *gin.Context) {
	req := &validate.TestModel{}
	err := validate.TransInit(req, "zh")
	//err := validate.InitTrans("zh", req)
	//req := &validate.TestModel{}
	//err := validate.DefaultGetValid(c, req)
	if err != nil {
		c.JSON(404, gin.H{
			"code": 2000,
			"err":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "success",
		"model": req,
	})
}

func main() {
	/*	model := validate.TestModel{
			CreateDate: time.Date(2022, time.July, 20, 0, 0, 0, 0, time.UTC),
			EndDate:    time.Date(2019, time.July, 20, 0, 0, 0, 0, time.UTC),
			CardId:     "100185201102304434",
			Phone:      "131234567890",
			PassportNo: "DE123456789",
		}
		validate.InitTrans("zh", model)*/
	route := gin.Default()
	route.POST("/test", test1)
	route.Run(":8099")

}
