// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRecoveryStrategy = "recovery_strategy"

// RecoveryStrategy mapped from table <recovery_strategy>
type RecoveryStrategy struct {
	RecoveryStrategyID     string   `gorm:"column:recovery_strategy_id;primaryKey" json:"recovery_strategy_id"` // ID
	RecoveryImplPackageID  *string  `gorm:"column:recovery_impl_package_id" json:"recovery_impl_package_id"`    // ID
	RecoveryImplID         *string  `gorm:"column:recovery_impl_id" json:"recovery_impl_id"`                    // ID
	ApplyAlertitems        *string  `gorm:"column:apply_alertitems" json:"apply_alertitems"`                    // 应用到的预警编码
	ApplyProjects          *string  `gorm:"column:apply_projects" json:"apply_projects"`                        // 应用到的项目
	ApplyIps               *string  `gorm:"column:apply_ips" json:"apply_ips"`                                  // 应用到的IP
	ApplySofts             *string  `gorm:"column:apply_softs" json:"apply_softs"`                              // 应用到的软件
	ServerOs               *string  `gorm:"column:server_os" json:"server_os"`                                  // 多个逗号分隔
	RecoveryStrategyName   *string  `gorm:"column:recovery_strategy_name" json:"recovery_strategy_name"`        // 方案名称
	SameRecoverSilenceSecs *float64 `gorm:"column:same_recover_silence_secs" json:"same_recover_silence_secs"`  // 相同对象的同样的自愈N秒内将忽略
	DelayRecoverSecs       *float64 `gorm:"column:delay_recover_secs" json:"delay_recover_secs"`                // 自愈执行延时N秒后开始
	IsValidate             *string  `gorm:"column:is_validate" json:"is_validate"`                              // 1、有效  0 无效
}

// TableName RecoveryStrategy's table name
func (*RecoveryStrategy) TableName() string {
	return TableNameRecoveryStrategy
}

func (*RecoveryStrategy) PrimaryKey() []string {
	return []string{"recovery_strategy_id"}
}