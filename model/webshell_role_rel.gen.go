// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWebshellRoleRel = "webshell_role_rel"

// WebshellRoleRel mapped from table <webshell_role_rel>
type WebshellRoleRel struct {
	WebshellUserRoleRelid string  `gorm:"column:webshell_user_role_relid;primaryKey" json:"webshell_user_role_relid"`
	UserGroupPermisionID  *string `gorm:"column:user_group_permision_id" json:"user_group_permision_id"` // UUID,主键
	RoleID                *string `gorm:"column:role_id" json:"role_id"`                                 // 角色ID
}

// TableName WebshellRoleRel's table name
func (*WebshellRoleRel) TableName() string {
	return TableNameWebshellRoleRel
}

func (*WebshellRoleRel) PrimaryKey() []string {
	return []string{"webshell_user_role_relid"}
}