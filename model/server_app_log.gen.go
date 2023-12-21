// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameServerAppLog = "server_app_log"

// ServerAppLog mapped from table <server_app_log>
type ServerAppLog struct {
	ServerAppLogID string     `gorm:"column:server_app_log_id;primaryKey" json:"server_app_log_id"`
	AppLogText     *string    `gorm:"column:app_log_text" json:"app_log_text"`
	OpTime         *time.Time `gorm:"column:op_time" json:"op_time"`
	OpUser         *string    `gorm:"column:op_user" json:"op_user"`
	BusiID         *string    `gorm:"column:busi_id" json:"busi_id"`
	CreateTime     *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime     *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments       *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName ServerAppLog's table name
func (*ServerAppLog) TableName() string {
	return TableNameServerAppLog
}

func (*ServerAppLog) PrimaryKey() []string {
	return []string{"server_app_log_id"}
}