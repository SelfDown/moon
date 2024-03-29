// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysDbuser = "sys_dbuser"

// SysDbuser mapped from table <sys_dbuser>
type SysDbuser struct {
	SysDbuserID      *string    `gorm:"column:sys_dbuser_id" json:"sys_dbuser_id"`
	DbuserValue      *string    `gorm:"column:dbuser_value" json:"dbuser_value"`
	DbuserText       *string    `gorm:"column:dbuser_text" json:"dbuser_text"`
	Note             *string    `gorm:"column:note" json:"note"`
	SysProjectTeamID *string    `gorm:"column:sys_project_team_id" json:"sys_project_team_id"`
	CreateTime       *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime       *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments         *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName SysDbuser's table name
func (*SysDbuser) TableName() string {
	return TableNameSysDbuser
}

func (*SysDbuser) PrimaryKey() []string {
	return []string{}
}