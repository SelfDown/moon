// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAutoCheckHostInfo = "auto_check_host_info"

// AutoCheckHostInfo mapped from table <auto_check_host_info>
type AutoCheckHostInfo struct {
	ID           int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ServerID     string     `gorm:"column:server_id;not null" json:"server_id"`                               // 主机id
	CheckItemsID *string    `gorm:"column:check_items_id" json:"check_items_id"`                              // 巡检项目集合,auto_check_items中的items_id是一对多的关系
	CreateTime   *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime   *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments     *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName AutoCheckHostInfo's table name
func (*AutoCheckHostInfo) TableName() string {
	return TableNameAutoCheckHostInfo
}

func (*AutoCheckHostInfo) PrimaryKey() []string {
	return []string{"id"}
}