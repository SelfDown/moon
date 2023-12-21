// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameServerOpsLog = "server_ops_log"

// ServerOpsLog mapped from table <server_ops_log>
type ServerOpsLog struct {
	OpsLogID      string     `gorm:"column:ops_log_id;primaryKey" json:"ops_log_id"`
	ServerID      *string    `gorm:"column:server_id" json:"server_id"`
	ActionName    string     `gorm:"column:action_name;not null" json:"action_name"`
	OsUser        string     `gorm:"column:os_user;not null" json:"os_user"`
	OperatorUser  *string    `gorm:"column:operator_user" json:"operator_user"`
	OperatorTime  *time.Time `gorm:"column:operator_time;not null;default:CURRENT_TIMESTAMP" json:"operator_time"`
	ActionNote    *string    `gorm:"column:action_note" json:"action_note"`
	ActionStatus  *string    `gorm:"column:action_status" json:"action_status"`
	SoftTypeName  *string    `gorm:"column:soft_type_name" json:"soft_type_name"`
	InstallSoftID *string    `gorm:"column:install_soft_id" json:"install_soft_id"`
}

// TableName ServerOpsLog's table name
func (*ServerOpsLog) TableName() string {
	return TableNameServerOpsLog
}

func (*ServerOpsLog) PrimaryKey() []string {
	return []string{"ops_log_id"}
}