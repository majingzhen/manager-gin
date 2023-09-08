package global

import (
	"time"
)

// GvaModel 基础模型
type GvaModel struct {
	Id         string     `json:"id" gorm:"id"` // 主键ID
	CreateBy   string     `json:"createBy" gorm:"create_by"`
	UpdateBy   string     `json:"updateBy" gorm:"update_by"`
	CreateTime *time.Time `json:"createTime" gorm:"create_time"` // 创建时间
	UpdateTime *time.Time `json:"updateTime" gorm:"update_time"` // 更新时间
	IsDel      int        `gorm:"index" json:"-"`                // 删除记录
}
