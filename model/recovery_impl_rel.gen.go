// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRecoveryImplRel = "recovery_impl_rel"

// RecoveryImplRel mapped from table <recovery_impl_rel>
type RecoveryImplRel struct {
	RecoveryImplRelID     string   `gorm:"column:recovery_impl_rel_id;primaryKey" json:"recovery_impl_rel_id"` // ID
	RecoveryImplPackageID *string  `gorm:"column:recovery_impl_package_id" json:"recovery_impl_package_id"`    // ID
	RecoveryImplID        *string  `gorm:"column:recovery_impl_id" json:"recovery_impl_id"`                    // ID
	Sort                  *float64 `gorm:"column:sort" json:"sort"`
}

// TableName RecoveryImplRel's table name
func (*RecoveryImplRel) TableName() string {
	return TableNameRecoveryImplRel
}

func (*RecoveryImplRel) PrimaryKey() []string {
	return []string{"recovery_impl_rel_id"}
}