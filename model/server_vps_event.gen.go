// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameServerVpsEvent = "server_vps_event"

// ServerVpsEvent mapped from table <server_vps_event>
type ServerVpsEvent struct {
	ServerVpsEventID int32      `gorm:"column:server_vps_event_id;primaryKey;autoIncrement:true" json:"server_vps_event_id"`
	CreateTime       *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime       *time.Time `gorm:"column:modify_time" json:"modify_time"`
	UserID           *string    `gorm:"column:user_id" json:"user_id"` // 操作用户
	/*
		事件操作细节
		比如 ：正在克隆虚拟机
	*/
	Comments *string `gorm:"column:comments" json:"comments"`
	/*
		事件类型与 server_vps_event_type,主键关联

	*/
	VpsEventTypeID *string `gorm:"column:vps_event_type_id" json:"vps_event_type_id"`
	ServerID       *string `gorm:"column:server_id" json:"server_id"`
	/*
		事件ID。
		比如用户申请创建5个虚拟机，那么这个算一次事件，我们将创建5个虚拟机的后台细节全部写入到表中，
		并且事件ID一样，这样前台可以根据事件ID，获取事件状态。
	*/
	EventID string `gorm:"column:event_id;not null" json:"event_id"`
}

// TableName ServerVpsEvent's table name
func (*ServerVpsEvent) TableName() string {
	return TableNameServerVpsEvent
}

func (*ServerVpsEvent) PrimaryKey() []string {
	return []string{"server_vps_event_id"}
}