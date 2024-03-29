// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUsersPermissions = "users_permissions"

// UsersPermissions mapped from table <users_permissions>
type UsersPermissions struct {
	ID           int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID       int32      `gorm:"column:user_id;not null" json:"user_id"`
	PermissionID int32      `gorm:"column:permission_id;not null" json:"permission_id"`
	CreateTime   *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime   *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments     *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName UsersPermissions's table name
func (*UsersPermissions) TableName() string {
	return TableNameUsersPermissions
}

func (*UsersPermissions) PrimaryKey() []string {
	return []string{"id"}
}