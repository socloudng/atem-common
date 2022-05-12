package base_model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BASE_MODEL[T uint | uint64 | int | int64 | uuid.UUID] struct {
	ID        T              `json:"id,string" form:"id" gorm:"primarykey"`            // 主键ID
	CreatedAt *time.Time     `json:"createTime" form:"createTime" gorm:"comment:创建时间"` // 创建时间
	UpdatedAt *time.Time     `json:"updateTime" form:"createTime" gorm:"comment:更新时间"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`                                   // 删除时间
}

/// 审计model
type AUDIT_MODEL struct {
	CreateBy int64 `json:"created_by" gorm:"comment:创建者"`
	UpdateBy int64 `json:"updated_by" gorm:"comment:修改者"`
}
