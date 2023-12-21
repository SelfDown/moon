// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDeptaskSpringbootConf = "deptask_springboot_conf"

// DeptaskSpringbootConf mapped from table <deptask_springboot_conf>
type DeptaskSpringbootConf struct {
	SpringbootConfID string     `gorm:"column:springboot_conf_id;primaryKey" json:"springboot_conf_id"`
	DepTaskID        *string    `gorm:"column:dep_task_id" json:"dep_task_id"`
	EnvConf          *string    `gorm:"column:env_conf" json:"env_conf"`
	Isvalid          *string    `gorm:"column:isvalid" json:"isvalid"` // 0 无效   1有效
	OpUser           *string    `gorm:"column:op_user" json:"op_user"`
	CreateTime       *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime       *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments         *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName DeptaskSpringbootConf's table name
func (*DeptaskSpringbootConf) TableName() string {
	return TableNameDeptaskSpringbootConf
}

func (*DeptaskSpringbootConf) PrimaryKey() []string {
	return []string{"springboot_conf_id"}
}