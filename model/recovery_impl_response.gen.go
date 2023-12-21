// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRecoveryImplResponse = "recovery_impl_response"

// RecoveryImplResponse mapped from table <recovery_impl_response>
type RecoveryImplResponse struct {
	RecoveryImplRespID     string     `gorm:"column:recovery_impl_resp_id;primaryKey" json:"recovery_impl_resp_id"` // ID
	RecoveryImplID         *string    `gorm:"column:recovery_impl_id" json:"recovery_impl_id"`                      // ID
	RecoveryImplPackageID  *string    `gorm:"column:recovery_impl_package_id" json:"recovery_impl_package_id"`      // ID
	AlertitemRecordID      *string    `gorm:"column:alertitem_record_id" json:"alertitem_record_id"`                // 预警消息ID
	RecoveryImplRespStatus *string    `gorm:"column:recovery_impl_resp_status" json:"recovery_impl_resp_status"`    // 1、成功 2、忽略 3、失败
	RecoveryImplRespRs     *string    `gorm:"column:recovery_impl_resp_rs" json:"recovery_impl_resp_rs"`            // 故障自愈执行结果
	RecoverImplRespRsType  *string    `gorm:"column:recover_impl_resp_rs_type" json:"recover_impl_resp_rs_type"`    // 结果类型 1、普通文档 2、xml 3、代码
	Note                   *string    `gorm:"column:note" json:"note"`                                              // 备注
	AddTime                *time.Time `gorm:"column:add_time;not null;default:CURRENT_TIMESTAMP" json:"add_time"`   // 添加时间
}

// TableName RecoveryImplResponse's table name
func (*RecoveryImplResponse) TableName() string {
	return TableNameRecoveryImplResponse
}

func (*RecoveryImplResponse) PrimaryKey() []string {
	return []string{"recovery_impl_resp_id"}
}