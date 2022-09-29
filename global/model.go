package global

import (
	"time"

	uuid "github.com/satori/go.uuid"

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
	// 适用于一些数据库自增性主键，数据库存储一般为数值，某些数据库可能不支持，比如cockroachdb并不会自增，这中ID在Mysql数据库中有更好的应用
	ID uint `gorm:"primarykey" json:"id,string" form:"id"` // 主键ID
	// UUID 通用的标准用户id，数据库存储为字符串在某些方面会更通用更好用
	UUID uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	// 用户名
	Account string `json:"account" gorm:"comment:用户登录名"`
	// 用户登录密码
	Password string `json:"-"  gorm:"comment:用户登录密码"`
	// 昵称
	Name string `json:"name" gorm:"default:系统用户;comment:用户昵称"`
	// 用户侧边主题
	SideMode string `json:"sideMode" gorm:"default:dark;comment:用户主题"`
	// 用户头像
	HeaderImg string `json:"headerImg" gorm:"default:https://c-ssl.dtstatic.com/uploads/item/201901/19/20190119105005_uJPTs.thumb.1000_0.jpeg;comment:用户头像"`
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
	// 用户手机号
	Phone string `json:"phone"  gorm:"comment:用户手机号"`
	// 用户邮箱
	Email string `json:"email"  gorm:"comment:用户邮箱"`

	// gorm会默认更新的字段

	CreatedAt time.Time      `json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`             // 删除时间
}
