// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDbscriptReqEvent = "dbscript_req_event"

// DbscriptReqEvent mapped from table <dbscript_req_event>
type DbscriptReqEvent struct {
	DbscriptEventID string     `gorm:"column:dbscript_event_id;primaryKey" json:"dbscript_event_id"`
	OpUser          *string    `gorm:"column:op_user" json:"op_user"`
	OpTime          *time.Time `gorm:"column:op_time" json:"op_time"`
	CheckFalg       *int32     `gorm:"column:check_falg" json:"check_falg"`
	CheckUser       *string    `gorm:"column:check_user" json:"check_user"`
	CheckTime       *time.Time `gorm:"column:check_time" json:"check_time"`
	ReqStatu        *int32     `gorm:"column:req_statu" json:"req_statu"` // 申请状态 0、已申请 1、已审核 2、已归档  3 用户下载
	HospitalCode    *string    `gorm:"column:hospital_code" json:"hospital_code"`
	CreateTime      *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime      *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments        *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName DbscriptReqEvent's table name
func (*DbscriptReqEvent) TableName() string {
	return TableNameDbscriptReqEvent
}

func (*DbscriptReqEvent) PrimaryKey() []string {
	return []string{"dbscript_event_id"}
}