// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameServerSSHEvent = "server_ssh_event"

// ServerSSHEvent mapped from table <server_ssh_event>
type ServerSSHEvent struct {
	ServerSSHEventID *string    `gorm:"column:server_ssh_event_id" json:"server_ssh_event_id"`
	IP               *string    `gorm:"column:ip" json:"ip"`
	SSHPort          *string    `gorm:"column:ssh_port" json:"ssh_port"`
	SSHUser          *string    `gorm:"column:ssh_user" json:"ssh_user"`
	Pwd              *string    `gorm:"column:pwd" json:"pwd"`
	Cmd              *string    `gorm:"column:cmd" json:"cmd"`
	Ret              *string    `gorm:"column:ret" json:"ret"`
	CreateTime       *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	Note             *string    `gorm:"column:note" json:"note"`
	TransactionID    *string    `gorm:"column:transaction_id" json:"transaction_id"`
	ModifyTime       *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments         *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName ServerSSHEvent's table name
func (*ServerSSHEvent) TableName() string {
	return TableNameServerSSHEvent
}

func (*ServerSSHEvent) PrimaryKey() []string {
	return []string{}
}