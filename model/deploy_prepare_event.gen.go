// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDeployPrepareEvent = "deploy_prepare_event"

// DeployPrepareEvent mapped from table <deploy_prepare_event>
type DeployPrepareEvent struct {
	DepPreID     string     `gorm:"column:dep_pre_id;primaryKey" json:"dep_pre_id"`
	OpUser       *string    `gorm:"column:op_user" json:"op_user"`             // 操作员
	OpTime       *time.Time `gorm:"column:op_time" json:"op_time"`             // 操作时间
	HospitalArea *string    `gorm:"column:hospital_area" json:"hospital_area"` // 院区
	DepRepID     *string    `gorm:"column:dep_rep_id" json:"dep_rep_id"`
	CreateTime   *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime   *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments     *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName DeployPrepareEvent's table name
func (*DeployPrepareEvent) TableName() string {
	return TableNameDeployPrepareEvent
}

func (*DeployPrepareEvent) PrimaryKey() []string {
	return []string{"dep_pre_id"}
}