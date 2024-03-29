// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameServerDbsyncDetail = "server_dbsync_detail"

// ServerDbsyncDetail mapped from table <server_dbsync_detail>
type ServerDbsyncDetail struct {
	ServerOracleSyncDetailID *string    `gorm:"column:server_oracle_sync_detail_id" json:"server_oracle_sync_detail_id"`
	CreateTime               *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	OpUser                   *string    `gorm:"column:op_user" json:"op_user"`
	Note                     *string    `gorm:"column:note" json:"note"`
	SrcServerEnvID           *string    `gorm:"column:src_server_env_id" json:"src_server_env_id"`                        // 数据来源数据库项目点
	DestServerEnvID          *string    `gorm:"column:dest_server_env_id" json:"dest_server_env_id"`                      // 数据同步目标数据库项目点
	SyncTime                 *time.Time `gorm:"column:sync_time" json:"sync_time"`                                        // 数据库同步时间
	ModifyTime               *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments                 *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName ServerDbsyncDetail's table name
func (*ServerDbsyncDetail) TableName() string {
	return TableNameServerDbsyncDetail
}

func (*ServerDbsyncDetail) PrimaryKey() []string {
	return []string{}
}