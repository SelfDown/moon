// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWebshellGroupRel = "webshell_group_rel"

// WebshellGroupRel mapped from table <webshell_group_rel>
type WebshellGroupRel struct {
	WebshellUserGroupRelID string  `gorm:"column:webshell_user_group_rel_id;primaryKey" json:"webshell_user_group_rel_id"`
	UserGroupPermisionID   *string `gorm:"column:user_group_permision_id" json:"user_group_permision_id"` // UUID,主键
	UserGroupCode          *string `gorm:"column:user_group_code" json:"user_group_code"`                 // user_group_code
}

// TableName WebshellGroupRel's table name
func (*WebshellGroupRel) TableName() string {
	return TableNameWebshellGroupRel
}

func (*WebshellGroupRel) PrimaryKey() []string {
	return []string{"webshell_user_group_rel_id"}
}