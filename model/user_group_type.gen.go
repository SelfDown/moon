// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserGroupType = "user_group_type"

// UserGroupType mapped from table <user_group_type>
type UserGroupType struct {
	Username  *string `gorm:"column:username" json:"username"`
	GroupType *string `gorm:"column:group_type" json:"group_type"`
	UserType  *string `gorm:"column:user_type" json:"user_type"`
	ID        string  `gorm:"column:id;primaryKey" json:"id"`
}

// TableName UserGroupType's table name
func (*UserGroupType) TableName() string {
	return TableNameUserGroupType
}

func (*UserGroupType) PrimaryKey() []string {
	return []string{"id"}
}