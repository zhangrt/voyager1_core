package global

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type GS_BASE_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"id,string" form:"id"` // 主键ID
	CreatedAt time.Time      `json:"createdAt" form:"createdAt"`            // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" form:"updatedAt"`            // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                        // 删除时间
}

// v0.2 针对基础业务做封装，被上层业务user结构体引用
type GS_BASE_USER struct {
	GS_BASE_MODEL
	// UUID
	UUID uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	// 用户名
	Username string `json:"userName" gorm:"comment:用户登录名"`
	// 昵称
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	// 权限ID
	AuthorityId string `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	// 部门Id
	DepartMentId string `json:"departmentId" gorm:"default:11111;comment:部门ID"`
	// 部门名称
	DepartMentName string `json:"departmentName" gorm:"default:综管部;comment:部门名称"`
	// 单位Id
	UnitId string `json:"unitId" gorm:"default:888;comment:单位ID"`
	// 单位名称
	UnitName string `json:"unitName" gorm:"default:888;comment:单位名称"`
}
