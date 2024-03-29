// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserRoleHospital = "user_role_hospital"

// UserRoleHospital mapped from table <user_role_hospital>
type UserRoleHospital struct {
	RoleID       *string    `gorm:"column:role_id" json:"role_id"`
	HospitalCode *string    `gorm:"column:hospital_code" json:"hospital_code"`
	IsDefault    *int32     `gorm:"column:is_default" json:"is_default"`                                      // 1 默认医院  0 非默认
	CreateTime   *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime   *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments     *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName UserRoleHospital's table name
func (*UserRoleHospital) TableName() string {
	return TableNameUserRoleHospital
}

func (*UserRoleHospital) PrimaryKey() []string {
	return []string{}
}