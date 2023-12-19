// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysHost = "sys_host"

// SysHost mapped from table <sys_host>
type SysHost struct {
	HostIPList     string     `gorm:"column:host_ip_list;primaryKey" json:"host_ip_list"` // 服务器IP
	HostName       *string    `gorm:"column:host_name" json:"host_name"`
	HostUserSuper  *string    `gorm:"column:host_user_super" json:"host_user_super"` // 主机超级管理员
	HostPwdSuper   *string    `gorm:"column:host_pwd_super" json:"host_pwd_super"`
	ProjectVersion *string    `gorm:"column:project_version" json:"project_version"`                            // 工程版本
	CreateTime     *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime     *time.Time `gorm:"column:modify_time" json:"modify_time"`
	Comments       *string    `gorm:"column:comments" json:"comments"`
	SysProjectID   *string    `gorm:"column:sys_project_id" json:"sys_project_id"`
	ProjectHome    *string    `gorm:"column:project_home" json:"project_home"` // 工程主目录
}

// TableName SysHost's table name
func (*SysHost) TableName() string {
	return TableNameSysHost
}

func (*SysHost) PrimaryKey() []string {
	return []string{"host_ip_list"}
}