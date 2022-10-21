package luna

import (
	"github.com/zhangrt/voyager1_core/global"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*CustomClaims, error) {
	token := c.Request.Header.Get(global.G_CONFIG.AUTHKey.Token)
	return GetUser(token)
}

// 解析token获取用户信息
func GetUser(token string) (*CustomClaims, error) {
	j := NewTOKEN()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.G_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) string {
	if claims, exists := c.Get(global.G_CONFIG.AUTHKey.User); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.ID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *gin.Context) []string {
	if claims, exists := c.Get(global.G_CONFIG.AUTHKey.User); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl.RoleIds
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.RoleIds
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get(global.G_CONFIG.AUTHKey.User); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}
