package global

import (
	"time"
)

// GvaModel 基础模型
type GvaModel struct {
	Id        string     `json:"id" gorm:"id,primary_key"` // 主键ID
	CreateBy  string     `json:"createBy" gorm:"create_by"`
	UpdateBy  string     `json:"updateBy" gorm:"update_by"`
	CreatedAt *time.Time `json:"createTime" gorm:"create_time"` // 创建时间
	UpdatedAt *time.Time `json:"updateTime" gorm:"update_time"` // 更新时间
	IsDel     int        `gorm:"index" json:"-"`                // 删除记录
}
