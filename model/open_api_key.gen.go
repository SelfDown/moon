// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOpenAPIKey = "open_api_key"

// OpenAPIKey mapped from table <open_api_key>
type OpenAPIKey struct {
	APIKey string `gorm:"column:api_key;not null" json:"api_key"` // api_key
	/*
		0、禁用
		1、启用
		2、已逻辑删除
	*/
	APIKeyStatus   *string    `gorm:"column:api_key_status" json:"api_key_status"`
	ExpirationTime *time.Time `gorm:"column:expiration_time" json:"expiration_time"` // 过期时间
	Userid         *string    `gorm:"column:userid" json:"userid"`                   // 令牌所属的用户
	/*
		key类型
		sys_code.type=open_api_key_type
	*/
	OpenAPIKeyType *string    `gorm:"column:open_api_key_type" json:"open_api_key_type"`
	CreateTime     *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime     *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments       *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
	OpUserid       *string    `gorm:"column:op_userid" json:"op_userid"`                                        // 操作人
	AppDomain      *string    `gorm:"column:app_domain" json:"app_domain"`                                      // 令牌应用域 ，见码表 type='app_domain'
	APIKeyID       int32      `gorm:"column:api_key_id;primaryKey;autoIncrement:true" json:"api_key_id"`
}

// TableName OpenAPIKey's table name
func (*OpenAPIKey) TableName() string {
	return TableNameOpenAPIKey
}

func (*OpenAPIKey) PrimaryKey() []string {
	return []string{"api_key_id"}
}