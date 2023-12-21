// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOpsSwDeployEvent = "ops_sw_deploy_event"

// OpsSwDeployEvent mapped from table <ops_sw_deploy_event>
type OpsSwDeployEvent struct {
	OpsDeployEventID  string     `gorm:"column:ops_deploy_event_id;not null" json:"ops_deploy_event_id"`           // 主键
	OpUserID          string     `gorm:"column:op_user_id;not null" json:"op_user_id"`                             // 事件操作人
	OpsSwDeployListID *string    `gorm:"column:ops_sw_deploy_list_id" json:"ops_sw_deploy_list_id"`                // ops_sw_deploy_list.ops_sw_deploy_list_id
	CreateTime        *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime        *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments          *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
	AnsibleExecDetail *string    `gorm:"column:ansible_exec_detail" json:"ansible_exec_detail"`                    // anisble执行任务的结果
	DeployDetail      *string    `gorm:"column:deploy_detail" json:"deploy_detail"`                                // 部署详细信息
	DeployStatus      *string    `gorm:"column:deploy_status" json:"deploy_status"`                                // 软件部署状态; sys_code.sys_type=deploy_event_status
	/*
		部署方式
		sys_code.sys_code_type='deploy_ops_sw_type'
	*/
	DeployOpsSwType *string `gorm:"column:deploy_ops_sw_type" json:"deploy_ops_sw_type"`
	ID              int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
}

// TableName OpsSwDeployEvent's table name
func (*OpsSwDeployEvent) TableName() string {
	return TableNameOpsSwDeployEvent
}

func (*OpsSwDeployEvent) PrimaryKey() []string {
	return []string{"id"}
}