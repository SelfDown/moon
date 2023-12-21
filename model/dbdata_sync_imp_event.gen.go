// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDbdataSyncImpEvent = "dbdata_sync_imp_event"

// DbdataSyncImpEvent mapped from table <dbdata_sync_imp_event>
type DbdataSyncImpEvent struct {
	DbsyncImpEventID string     `gorm:"column:dbsync_imp_event_id;not null" json:"dbsync_imp_event_id"`
	InstallSoftID    *string    `gorm:"column:install_soft_id" json:"install_soft_id"`
	SoftUserID       *string    `gorm:"column:soft_user_id" json:"soft_user_id"`
	DbsyncExpEventID *string    `gorm:"column:dbsync_exp_event_id" json:"dbsync_exp_event_id"`
	TaskTitle        *string    `gorm:"column:task_title" json:"task_title"`
	StepIndex        *string    `gorm:"column:step_index" json:"step_index"`
	SyncStartTime    *time.Time `gorm:"column:sync_start_time;default:CURRENT_TIMESTAMP" json:"sync_start_time"`
	SyncEndTime      *time.Time `gorm:"column:sync_end_time;default:CURRENT_TIMESTAMP" json:"sync_end_time"`
	SyncCurrentStep  *string    `gorm:"column:sync_current_step" json:"sync_current_step"`
	SyncLog          *string    `gorm:"column:sync_log" json:"sync_log"`
	SyncScript       *string    `gorm:"column:sync_script" json:"sync_script"`
	AddTime          *time.Time `gorm:"column:add_time;default:CURRENT_TIMESTAMP" json:"add_time"`
	OpUser           *string    `gorm:"column:op_user" json:"op_user"`
	Note             *string    `gorm:"column:note" json:"note"`
	SyncStatu        *string    `gorm:"column:sync_statu" json:"sync_statu"`
	DirectoryPath    *string    `gorm:"column:directory_path" json:"directory_path"`
	OpPercent        *string    `gorm:"column:op_percent" json:"op_percent"`
	DbsyncSrcFile    *string    `gorm:"column:dbsync_src_file" json:"dbsync_src_file"`
	CanEditor        *string    `gorm:"column:can_editor" json:"can_editor"`
	StransactionIdx  *string    `gorm:"column:stransaction_idx" json:"stransaction_idx"`
}

// TableName DbdataSyncImpEvent's table name
func (*DbdataSyncImpEvent) TableName() string {
	return TableNameDbdataSyncImpEvent
}

func (*DbdataSyncImpEvent) PrimaryKey() []string {
	return []string{}
}