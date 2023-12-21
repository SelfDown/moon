// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserGroupRel = "user_group_rel"

// UserGroupRel mapped from table <user_group_rel>
type UserGroupRel struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID      *string    `gorm:"column:user_id" json:"user_id"`
	CreateTime  *time.Time `gorm:"column:create_time" json:"create_time"`
	ModifyTime  *time.Time `gorm:"column:modify_time" json:"modify_time"`
	Comments    *string    `gorm:"column:comments" json:"comments"`
	UserGroupID *string    `gorm:"column:user_group_id" json:"user_group_id"`
}

// TableName UserGroupRel's table name
func (*UserGroupRel) TableName() string {
	return TableNameUserGroupRel
}

func (*UserGroupRel) PrimaryKey() []string {
	return []string{"id"}
}