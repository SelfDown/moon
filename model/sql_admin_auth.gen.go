// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSqlAdminAuth = "sql_admin_auth"

// SqlAdminAuth mapped from table <sql_admin_auth>
type SqlAdminAuth struct {
	SqlAdminAuthID string `gorm:"column:sql_admin_auth_id;primaryKey" json:"sql_admin_auth_id"`
	/*
		环境类型
		生产
		非生产
	*/
	EnvType        *string    `gorm:"column:env_type" json:"env_type"`
	SqlAdminRoleID *string    `gorm:"column:sql_admin_role_id" json:"sql_admin_role_id"`                        // 关联 sql_admin_role.sql_admin_role_id
	AuthBy         *string    `gorm:"column:auth_by" json:"auth_by"`                                            // 创建这条记录的操作人 user_id
	CreateTime     *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime     *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments       *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
	UserID         string     `gorm:"column:user_id;not null" json:"user_id"`                                   // 被授权人userid
}

// TableName SqlAdminAuth's table name
func (*SqlAdminAuth) TableName() string {
	return TableNameSqlAdminAuth
}

func (*SqlAdminAuth) PrimaryKey() []string {
	return []string{"sql_admin_auth_id"}
}