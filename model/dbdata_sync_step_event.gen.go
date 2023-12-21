// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDbdataSyncStepEvent = "dbdata_sync_step_event"

// DbdataSyncStepEvent mapped from table <dbdata_sync_step_event>
type DbdataSyncStepEvent struct {
	DbdataSyncStepEventID string     `gorm:"column:dbdata_sync_step_event_id;not null" json:"dbdata_sync_step_event_id"`
	DbsyncExpEventID      *string    `gorm:"column:dbsync_exp_event_id" json:"dbsync_exp_event_id"`
	DbsyncImpEventID      *string    `gorm:"column:dbsync_imp_event_id" json:"dbsync_imp_event_id"`
	StepScript            *string    `gorm:"column:step_script" json:"step_script"`
	StepName              *string    `gorm:"column:step_name" json:"step_name"`
	StepSyncLog           *string    `gorm:"column:step_sync_log" json:"step_sync_log"`
	StepReqStr            *string    `gorm:"column:step_req_str" json:"step_req_str"`
	StepStatu             *string    `gorm:"column:step_statu" json:"step_statu"`
	AddTime               *time.Time `gorm:"column:add_time;default:CURRENT_TIMESTAMP" json:"add_time"`
	OpUser                string     `gorm:"column:op_user;not null" json:"op_user"`
	Note                  *string    `gorm:"column:note" json:"note"`
}

// TableName DbdataSyncStepEvent's table name
func (*DbdataSyncStepEvent) TableName() string {
	return TableNameDbdataSyncStepEvent
}

func (*DbdataSyncStepEvent) PrimaryKey() []string {
	return []string{}
}