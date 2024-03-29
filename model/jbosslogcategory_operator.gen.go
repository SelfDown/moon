// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameJbosslogcategoryOperator = "jbosslogcategory_operator"

// JbosslogcategoryOperator mapped from table <jbosslogcategory_operator>
type JbosslogcategoryOperator struct {
	EventID       string     `gorm:"column:event_id;primaryKey" json:"event_id"`
	InstallSoftID *string    `gorm:"column:install_soft_id" json:"install_soft_id"`
	ServerIP      *string    `gorm:"column:server_ip" json:"server_ip"`
	ServerPort    *string    `gorm:"column:server_port" json:"server_port"`
	ExpiryTime    *time.Time `gorm:"column:expiry_time;not null;default:CURRENT_TIMESTAMP" json:"expiry_time"` // 过期时间
	LastExeRs     *string    `gorm:"column:last_exe_rs" json:"last_exe_rs"`                                    // 最近执行结果日志
	CategoryName  *string    `gorm:"column:category_name" json:"category_name"`
	CategoryLevel *string    `gorm:"column:category_level" json:"category_level"`
	Status        *string    `gorm:"column:status" json:"status"` // 1、新入  2、已撤销
	OpUser        *string    `gorm:"column:op_user" json:"op_user"`
	RemoveTime    time.Time  `gorm:"column:remove_time;not null" json:"remove_time"`
	AddTime       *time.Time `gorm:"column:add_time;not null;default:CURRENT_TIMESTAMP" json:"add_time"` // 记录创建时间（数据库自动写入）
	SoftType      *string    `gorm:"column:soft_type" json:"soft_type"`
	Comments      *string    `gorm:"column:comments" json:"comments"` // 备注说明
	Project       *string    `gorm:"column:project" json:"project"`
}

// TableName JbosslogcategoryOperator's table name
func (*JbosslogcategoryOperator) TableName() string {
	return TableNameJbosslogcategoryOperator
}

func (*JbosslogcategoryOperator) PrimaryKey() []string {
	return []string{"event_id"}
}
