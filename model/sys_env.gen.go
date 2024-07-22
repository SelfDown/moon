// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysEnv = "sys_env"

// SysEnv mapped from table <sys_env>
type SysEnv struct {
	EnvCode    *string `gorm:"column:env_code;primaryKey" json:"env_code"`
	EnvName    string  `gorm:"column:env_name;not null" json:"env_name"`
	EnvNote    *string `gorm:"column:env_note" json:"env_note"`
	CreateTime *string `gorm:"column:create_time" json:"create_time"`
	ModifyTime *string `gorm:"column:modify_time" json:"modify_time"`
	Comments   *string `gorm:"column:comments" json:"comments"`
}

// TableName SysEnv's table name
func (*SysEnv) TableName() string {
	return TableNameSysEnv
}

func (*SysEnv) PrimaryKey() []string {
	return []string{"env_code"}
}