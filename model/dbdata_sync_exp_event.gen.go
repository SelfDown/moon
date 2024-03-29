// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDbdataSyncExpEvent = "dbdata_sync_exp_event"

// DbdataSyncExpEvent mapped from table <dbdata_sync_exp_event>
type DbdataSyncExpEvent struct {
	DbsyncExpEventID string     `gorm:"column:dbsync_exp_event_id;primaryKey" json:"dbsync_exp_event_id"`
	SoftUserID       *string    `gorm:"column:soft_user_id" json:"soft_user_id"`
	InstallSoftID    *string    `gorm:"column:install_soft_id" json:"install_soft_id"`
	TaskTitle        *string    `gorm:"column:task_title" json:"task_title"`
	SyncStartTime    *time.Time `gorm:"column:sync_start_time;default:CURRENT_TIMESTAMP" json:"sync_start_time"`
	SyncEndTime      *time.Time `gorm:"column:sync_end_time;default:CURRENT_TIMESTAMP" json:"sync_end_time"`
	StepIndex        *string    `gorm:"column:step_index" json:"step_index"`
	SyncCurrentStep  *string    `gorm:"column:sync_current_step" json:"sync_current_step"`
	SyncLog          *string    `gorm:"column:sync_log" json:"sync_log"`
	SyncScript       *string    `gorm:"column:sync_script" json:"sync_script"`
	AddTime          *time.Time `gorm:"column:add_time;default:CURRENT_TIMESTAMP" json:"add_time"`
	OpUser           string     `gorm:"column:op_user;not null" json:"op_user"`
	Note             *string    `gorm:"column:note" json:"note"`
	SyncStatu        *string    `gorm:"column:sync_statu" json:"sync_statu"`
	UploadFileName   *string    `gorm:"column:upload_file_name" json:"upload_file_name"`
	UploadFilePath   *string    `gorm:"column:upload_file_path" json:"upload_file_path"`
	DirectoryPath    *string    `gorm:"column:directory_path" json:"directory_path"`
	OpPercent        *string    `gorm:"column:op_percent" json:"op_percent"`
	CanEditor        *string    `gorm:"column:can_editor" json:"can_editor"`
	StransactionIdx  *string    `gorm:"column:stransaction_idx" json:"stransaction_idx"`
	ExpOraVersion    *string    `gorm:"column:exp_ora_version" json:"exp_ora_version"`
}

// TableName DbdataSyncExpEvent's table name
func (*DbdataSyncExpEvent) TableName() string {
	return TableNameDbdataSyncExpEvent
}

func (*DbdataSyncExpEvent) PrimaryKey() []string {
	return []string{"dbsync_exp_event_id"}
}