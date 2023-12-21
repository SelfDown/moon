// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMsgSendToken = "msg_send_token"

// MsgSendToken mapped from table <msg_send_token>
type MsgSendToken struct {
	MsgSendTokenID string     `gorm:"column:msg_send_token_id;primaryKey" json:"msg_send_token_id"` // ID
	Token          *string    `gorm:"column:token" json:"token"`                                    // 令牌值
	TimeBegin      *time.Time `gorm:"column:time_begin" json:"time_begin"`                          // 生效时间-开始
	TimeEnd        *time.Time `gorm:"column:time_end" json:"time_end"`                              // 生效时间-结束
	AddTime        *time.Time `gorm:"column:add_time" json:"add_time"`                              // 添加时间
	AddUser        *string    `gorm:"column:add_user" json:"add_user"`                              // 添加用户
	Note           *string    `gorm:"column:note" json:"note"`                                      // 备注
	TokenName      *string    `gorm:"column:token_name" json:"token_name"`
}

// TableName MsgSendToken's table name
func (*MsgSendToken) TableName() string {
	return TableNameMsgSendToken
}

func (*MsgSendToken) PrimaryKey() []string {
	return []string{"msg_send_token_id"}
}