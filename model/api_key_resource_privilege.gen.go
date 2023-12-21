// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAPIKeyResourcePrivilege = "api_key_resource_privilege"

// APIKeyResourcePrivilege mapped from table <api_key_resource_privilege>
type APIKeyResourcePrivilege struct {
	ID                   int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	APIKeyResourceTypeID *int32     `gorm:"column:api_key_resource_type_id" json:"api_key_resource_type_id"`          // 对应api_key_resource _type 的id
	RoleID               *string    `gorm:"column:role_id" json:"role_id"`                                            // 附属角色id
	PrivilegeField       *string    `gorm:"column:privilege_field" json:"privilege_field"`                            // 特权字段，用"|"分隔
	CreateTime           *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime           *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments             *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName APIKeyResourcePrivilege's table name
func (*APIKeyResourcePrivilege) TableName() string {
	return TableNameAPIKeyResourcePrivilege
}

func (*APIKeyResourcePrivilege) PrimaryKey() []string {
	return []string{"id"}
}