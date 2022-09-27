package util

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangrt/voyager1_core/util/validate"
	"net/http"
)

func test1(c *gin.Context) {
	req := &validate.TestModel{}
	err := validate.TransInit(req, "zh")
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
	route := gin.Default()
	route.POST("/test", test1)
	route.Run(":8099")
}
