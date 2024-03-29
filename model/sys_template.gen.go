// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysTemplate = "sys_template"

// SysTemplate mapped from table <sys_template>
type SysTemplate struct {
	SysTemplateID   string     `gorm:"column:sys_template_id;primaryKey" json:"sys_template_id"`
	TemplateContent *string    `gorm:"column:template_content" json:"template_content"`
	TemplateName    *string    `gorm:"column:template_name" json:"template_name"`
	Comments        *string    `gorm:"column:comments" json:"comments"`
	CreateTime      *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime      *time.Time `gorm:"column:modify_time" json:"modify_time"`
	UserID          *string    `gorm:"column:user_id" json:"user_id"`
}

// TableName SysTemplate's table name
func (*SysTemplate) TableName() string {
	return TableNameSysTemplate
}

func (*SysTemplate) PrimaryKey() []string {
	return []string{"sys_template_id"}
}