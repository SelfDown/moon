// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRecoveryImpl = "recovery_impl"

// RecoveryImpl mapped from table <recovery_impl>
type RecoveryImpl struct {
	RecoveryImplID   string  `gorm:"column:recovery_impl_id;primaryKey" json:"recovery_impl_id"` // ID
	RecoverService   *string `gorm:"column:recover_service" json:"recover_service"`              // 服务方案
	RecoverParams    *string `gorm:"column:recover_params" json:"recover_params"`                // 内部方法参数，以Json字符串表达
	RecoveryImplName *string `gorm:"column:recovery_impl_name" json:"recovery_impl_name"`        // 恢复实现方法名称
	Note             *string `gorm:"column:note" json:"note"`                                    // 备注
	IsValidate       *string `gorm:"column:is_validate" json:"is_validate"`                      // 1、有效  0 无效
	ResultRecType    *string `gorm:"column:result_rec_type" json:"result_rec_type"`              // 1、默认仅仅记录返回，2发送消息+记录
}

// TableName RecoveryImpl's table name
func (*RecoveryImpl) TableName() string {
	return TableNameRecoveryImpl
}

func (*RecoveryImpl) PrimaryKey() []string {
	return []string{"recovery_impl_id"}
}