// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysTokenUserGroup = "sys_token_user_group"

// SysTokenUserGroup mapped from table <sys_token_user_group>
type SysTokenUserGroup struct {
	SysTokenUserGroupID string  `gorm:"column:sys_token_user_group_id;primaryKey" json:"sys_token_user_group_id"`
	SysTokenID          *string `gorm:"column:sys_token_id" json:"sys_token_id"`
	UserGroupID         *string `gorm:"column:user_group_id" json:"user_group_id"`
}

// TableName SysTokenUserGroup's table name
func (*SysTokenUserGroup) TableName() string {
	return TableNameSysTokenUserGroup
}

func (*SysTokenUserGroup) PrimaryKey() []string {
	return []string{"sys_token_user_group_id"}
}