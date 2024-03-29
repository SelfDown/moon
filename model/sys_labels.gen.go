// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysLabels = "sys_labels"

// SysLabels mapped from table <sys_labels>
type SysLabels struct {
	LabelID      string     `gorm:"column:label_id;primaryKey" json:"label_id"`
	LabelGroup   string     `gorm:"column:label_group;not null" json:"label_group"`
	LabelName    string     `gorm:"column:label_name;not null" json:"label_name"`
	OperatorUser string     `gorm:"column:operator_user;not null" json:"operator_user"`
	OperatorTime *time.Time `gorm:"column:operator_time;not null;default:CURRENT_TIMESTAMP" json:"operator_time"`
}

// TableName SysLabels's table name
func (*SysLabels) TableName() string {
	return TableNameSysLabels
}

func (*SysLabels) PrimaryKey() []string {
	return []string{"label_id"}
}