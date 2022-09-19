package luna

import (
	"github.com/xyy277/gallery/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Enforce(c *gin.Context) (bool, error) {
	waitUse, _ := GetClaims(c)
	// 获取请求的PATH
	obj := c.Request.URL.Path
	// 获取请求方法
	act := c.Request.Method
	// 获取用户的角色
	sub := waitUse.AuthorityId
	auth := NewCasbin()
	e := auth.Casbin()
	// 判断策略中是否存在
	success, err := e.Enforce(sub, obj, act)

	return success, err
}

func LoadAll() {
	var data []string
	err := global.G_DB.Model(&JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.G_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
