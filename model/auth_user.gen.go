// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAuthUser = "auth_user"

// AuthUser mapped from table <auth_user>
type AuthUser struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Password    string     `gorm:"column:password;not null" json:"password"`
	LastLogin   *time.Time `gorm:"column:last_login" json:"last_login"`
	IsSuperuser bool       `gorm:"column:is_superuser;not null" json:"is_superuser"`
	Email       string     `gorm:"column:email;not null" json:"email"`
	IsStaff     bool       `gorm:"column:is_staff;not null" json:"is_staff"`
	IsActive    bool       `gorm:"column:is_active;not null" json:"is_active"`
	DateJoined  time.Time  `gorm:"column:date_joined;not null" json:"date_joined"`
	UserTypeID  *int32     `gorm:"column:user_type_id" json:"user_type_id"`
	CreateTime  *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime  *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录创建时间（数据库自动写入）
	Comments    *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName AuthUser's table name
func (*AuthUser) TableName() string {
	return TableNameAuthUser
}

func (*AuthUser) PrimaryKey() []string {
	return []string{"id"}
}