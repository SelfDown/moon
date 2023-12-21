// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysCommonConfig = "sys_common_config"

// SysCommonConfig mapped from table <sys_common_config>
type SysCommonConfig struct {
	SysCommonConfigID string  `gorm:"column:sys_common_config_id;primaryKey" json:"sys_common_config_id"`
	ProjectID         string  `gorm:"column:project_id;not null" json:"project_id"` // 项目ID
	Key               string  `gorm:"column:key;not null" json:"key"`               // 键
	Value             *string `gorm:"column:value" json:"value"`                    // 值
	Type              *string `gorm:"column:type" json:"type"`                      // 类型 json或者字符串
	Comment           *string `gorm:"column:comment" json:"comment"`                // 备注
}

// TableName SysCommonConfig's table name
func (*SysCommonConfig) TableName() string {
	return TableNameSysCommonConfig
}

func (*SysCommonConfig) PrimaryKey() []string {
	return []string{"sys_common_config_id"}
}