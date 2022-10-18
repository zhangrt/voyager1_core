package global

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type GS_BASE_MODEL_ID_UINT struct {
	ID            uint           `gorm:"primarykey" json:"id,string" form:"id"` // 主键ID, uint json 需要指定转为string
	CreatorId     string         `json:"creator_id"  gorm:"comment:创建人id"`
	Creator       string         `json:"creator"  gorm:"comment:创建人"`
	LastUpdaterId string         `json:"last_update_id"  gorm:"comment:更新人id"`
	LastUpdater   string         `json:"last_updater"  gorm:"comment:更新人"`
	CreatedAt     time.Time      `json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt     time.Time      `json:"updatedAt" form:"updatedAt"` // 更新时间
	Deleted       int            `json:"deleted"  gorm:"column:deleted;size:1;comment:删除标记"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type GS_BASE_MODEL_ID_STRING struct {
	ID            string         `gorm:"primarykey" json:"id" form:"id"` // 主键ID
	CreatorId     string         `json:"creator_id"  gorm:"comment:创建人id"`
	Creator       string         `json:"creator"  gorm:"comment:创建人"`
	LastUpdaterId string         `json:"last_update_id"  gorm:"comment:更新人id"`
	LastUpdater   string         `json:"last_updater"  gorm:"comment:更新人"`
	CreatedAt     time.Time      `json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt     time.Time      `json:"updatedAt" form:"updatedAt"` // 更新时间
	Deleted       int            `json:"deleted"  gorm:"column:deleted;size:1;comment:删除标记"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type GS_BASE_MODEL_ID_NO_PRIMARY struct {
	Id            string         `json:"id" form:"id"` // ID
	CreatorId     string         `json:"creator_id"  gorm:"comment:创建人id"`
	Creator       string         `json:"creator"  gorm:"comment:创建人"`
	LastUpdaterId string         `json:"last_update_id"  gorm:"comment:更新人id"`
	LastUpdater   string         `json:"last_updater"  gorm:"comment:更新人"`
	CreatedAt     time.Time      `json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt     time.Time      `json:"updatedAt" form:"updatedAt"` // 更新时间
	Deleted       int            `json:"deleted"  gorm:"column:deleted;size:1;comment:删除标记"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type GS_BASE_MODEL_ID_NONE struct {
	CreatorId     string         `json:"creator_id"  gorm:"comment:创建人id"`
	Creator       string         `json:"creator"  gorm:"comment:创建人"`
	LastUpdaterId string         `json:"last_update_id"  gorm:"comment:更新人id"`
	LastUpdater   string         `json:"last_updater"  gorm:"comment:更新人"`
	CreatedAt     time.Time      `json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt     time.Time      `json:"updatedAt" form:"updatedAt"` // 更新时间
	Deleted       int            `json:"deleted"  gorm:"column:deleted;size:1;comment:删除标记"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

// v0.2 针对基础业务做封装，被上层业务user结构体引用
type GS_BASE_USER struct {
	// 适用于一些数据库自增性主键，数据库存储一般为数值，某些数据库可能不支持，比如cockroachdb并不会自增，这中ID在Mysql数据库中有更好的应用
	// ID string `gorm:"primarykey" json:"id,string" form:"id"` // 主键ID
	ID uuid.UUID `gorm:"primarykey" json:"id" form:"id"` // 主键ID
	// UUID 通用的标准用户id，数据库存储为字符串在某些方面会更通用更好用
	// UUID uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	// 用户名
	Account string `json:"account" gorm:"comment:用户登录名"`
	// 用户登录密码
	Password string `json:"-"  gorm:"comment:用户登录密码"`
	// 昵称
	Name     string `json:"name" gorm:"default:系统用户;comment:用户昵称"`
	Age      string `json:"age" gorm:"comment:年龄"`
	Gender   string `json:"gender" gorm:"comment:性别"`
	SerialNo string `json:"serialNo" gorm:"comment:排序字段"`
	// 用户侧边主题
	// SideMode string `json:"sideMode" gorm:"default:dark;comment:用户主题"`
	// 用户头像
	Avatar string `json:"avatar" gorm:"default:https://c-ssl.dtstatic.com/uploads/item/201901/19/20190119105005_uJPTs.thumb.1000_0.jpeg;comment:用户头像"`
	// 用户手机号
	Phone string `json:"phone"  gorm:"comment:用户手机号"`
	// 用户邮箱
	Email string `json:"email"  gorm:"comment:用户邮箱"`
	// 用户锁定
	Locked bool `json:"locked"  gorm:"column:locked;size:1;comment:用户锁定"`
	// 锁定时间
	LockTime time.Time `json:"lockTime"  gorm:"column:lock_time;comment:用户锁定时间"`
	// 最后一次登录时间
	LastLoginTime time.Time `json:"lastLoginTime"  gorm:"comment:最后一次登录时间"`
	// 描述
	Description string `json:"description"  gorm:"comment:用户描述"`
	// 创建人id
	CreatorId string `json:"creator_id"  gorm:"comment:创建人id"`
	// 创建人
	Creator string `json:"creator"  gorm:"comment:创建人"`
	// 创建时间
	LastUpdaterId string         `json:"last_update_id"  gorm:"comment:更新人id"`
	LastUpdater   string         `json:"last_updater"  gorm:"comment:更新人"`
	CreatedAt     time.Time      `json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt     time.Time      `json:"updatedAt" form:"updatedAt"` // 更新时间
	Deleted       int            `json:"deleted"  gorm:"column:deleted;size:1;comment:删除标记"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
