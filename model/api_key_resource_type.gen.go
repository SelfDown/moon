// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAPIKeyResourceType = "api_key_resource_type"

// APIKeyResourceType mapped from table <api_key_resource_type>
type APIKeyResourceType struct {
	ID                 int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ResourceType       string     `gorm:"column:resource_type;not null" json:"resource_type"`                       // 资源类型，对应sys_code 的api_key 类型
	ResourceDescribe   *string    `gorm:"column:resource_describe" json:"resource_describe"`                        // 资源描述
	ResourceTable      string     `gorm:"column:resource_table;not null" json:"resource_table"`                     // 资源对应的表
	ResourcePrimaryKey *string    `gorm:"column:resource_primary_key" json:"resource_primary_key"`                  // 资源表对应的主键
	CreateTime         *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime         *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments           *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName APIKeyResourceType's table name
func (*APIKeyResourceType) TableName() string {
	return TableNameAPIKeyResourceType
}
