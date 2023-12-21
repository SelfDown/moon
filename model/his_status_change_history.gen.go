// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHisStatusChangeHistory = "his_status_change_history"

// HisStatusChangeHistory mapped from table <his_status_change_history>
type HisStatusChangeHistory struct {
	HisStatusChangeHistoryID string  `gorm:"column:his_status_change_history_id;primaryKey" json:"his_status_change_history_id"` // 主键，UUID
	HisIssueRecordID         *string `gorm:"column:his_issue_record_id" json:"his_issue_record_id"`
	RecordStatus             *string `gorm:"column:record_status" json:"record_status"`
	OpUser                   *string `gorm:"column:op_user" json:"op_user"` // 状态修改人
	Receiver                 *string `gorm:"column:receiver" json:"receiver"`
	CreateTime               *string `gorm:"column:create_time" json:"create_time"`
	ModifyTime               *string `gorm:"column:modify_time" json:"modify_time"`
	Comments                 *string `gorm:"column:comments" json:"comments"` // 备注说明,记录调整备注说明
	ReceiverTel              *string `gorm:"column:receiver_tel" json:"receiver_tel"`
	IsOuterUser              *string `gorm:"column:is_outer_user" json:"is_outer_user"`
	HisConfirmWorkCode       *string `gorm:"column:his_confirm_work_code" json:"his_confirm_work_code"`
	HisConfirmPhone          *string `gorm:"column:his_confirm_phone" json:"his_confirm_phone"`
	HisConfirmUser           *string `gorm:"column:his_confirm_user" json:"his_confirm_user"`
}

// TableName HisStatusChangeHistory's table name
func (*HisStatusChangeHistory) TableName() string {
	return TableNameHisStatusChangeHistory
}

func (*HisStatusChangeHistory) PrimaryKey() []string {
	return []string{"his_status_change_history_id"}
}