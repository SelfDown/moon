// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAuthUserCopy = "auth_user_copy"

// AuthUserCopy mapped from table <auth_user_copy>
type AuthUserCopy struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Password    string     `gorm:"column:password;not null" json:"password"`
	LastLogin   *time.Time `gorm:"column:last_login" json:"last_login"`
	IsSuperuser bool       `gorm:"column:is_superuser;not null" json:"is_superuser"`
	Email       string     `gorm:"column:email;not null" json:"email"`
	IsStaff     bool       `gorm:"column:is_staff;not null" json:"is_staff"`
	IsActive    bool       `gorm:"column:is_active;not null" json:"is_active"`
	DateJoined  time.Time  `gorm:"column:date_joined;not null" json:"date_joined"`
	UserTypeID  *int32     `gorm:"column:user_type_id" json:"user_type_id"`
	CreateTime  *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	ModifyTime  *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"`
	Comments    *string    `gorm:"column:comments" json:"comments"`
}

// TableName AuthUserCopy's table name
func (*AuthUserCopy) TableName() string {
	return TableNameAuthUserCopy
}

func (*AuthUserCopy) PrimaryKey() []string {
	return []string{"id"}
}