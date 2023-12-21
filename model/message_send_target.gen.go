// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMessageSendTarget = "message_send_target"

// MessageSendTarget mapped from table <message_send_target>
type MessageSendTarget struct {
	MessageTargetID string  `gorm:"column:message_target_id;primaryKey" json:"message_target_id"`
	TargetID        *string `gorm:"column:target_id" json:"target_id"` // 发送对象ID（比如某个群的ID）
	/*
		对象类型
		0 邮件
		1 QQ
		2 微信

	*/
	TargetType   *string    `gorm:"column:target_type" json:"target_type"`
	TargetDesc   *string    `gorm:"column:target_desc" json:"target_desc"`                                    // 对象描述（比如：某个群的名字）
	SysProjectID *string    `gorm:"column:sys_project_id" json:"sys_project_id"`                              // 所属项目
	Statu        *string    `gorm:"column:statu" json:"statu"`                                                // 1 启用 0 禁用
	CreateTime   *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime   *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments     *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName MessageSendTarget's table name
func (*MessageSendTarget) TableName() string {
	return TableNameMessageSendTarget
}

func (*MessageSendTarget) PrimaryKey() []string {
	return []string{"message_target_id"}
}