// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMsgSendRelChannel = "msg_send_rel_channel"

// MsgSendRelChannel mapped from table <msg_send_rel_channel>
type MsgSendRelChannel struct {
	SendRelChannelID string  `gorm:"column:send_rel_channel_id;primaryKey" json:"send_rel_channel_id"`
	MsgSendRuleID    *string `gorm:"column:msg_send_rule_id" json:"msg_send_rule_id"`
	ChannelID        *string `gorm:"column:channel_id" json:"channel_id"`
	ChannelGroupID   *string `gorm:"column:channel_group_id" json:"channel_group_id"`
	RelChannelType   *string `gorm:"column:rel_channel_type" json:"rel_channel_type"`
}

// TableName MsgSendRelChannel's table name
func (*MsgSendRelChannel) TableName() string {
	return TableNameMsgSendRelChannel
}

func (*MsgSendRelChannel) PrimaryKey() []string {
	return []string{"send_rel_channel_id"}
}