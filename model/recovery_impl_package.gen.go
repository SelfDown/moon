// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRecoveryImplPackage = "recovery_impl_package"

// RecoveryImplPackage mapped from table <recovery_impl_package>
type RecoveryImplPackage struct {
	RecoveryImplPackageID   string  `gorm:"column:recovery_impl_package_id;primaryKey" json:"recovery_impl_package_id"` // ID
	RecoveryImplPackageName *string `gorm:"column:recovery_impl_package_name" json:"recovery_impl_package_name"`        // 恢复实现方法名称
	Note                    *string `gorm:"column:note" json:"note"`                                                    // 备注
}

// TableName RecoveryImplPackage's table name
func (*RecoveryImplPackage) TableName() string {
	return TableNameRecoveryImplPackage
}

func (*RecoveryImplPackage) PrimaryKey() []string {
	return []string{"recovery_impl_package_id"}
}