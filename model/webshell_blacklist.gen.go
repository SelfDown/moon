// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameWebshellBlacklist = "webshell_blacklist"

// WebshellBlacklist mapped from table <webshell_blacklist>
type WebshellBlacklist struct {
	WebshellBlacklistID string     `gorm:"column:webshell_blacklist_id;not null" json:"webshell_blacklist_id"`
	AddTime             *time.Time `gorm:"column:add_time" json:"add_time"`
	Comments            *string    `gorm:"column:comments" json:"comments"`
	OpUser              *string    `gorm:"column:op_user" json:"op_user"`
	Action              *string    `gorm:"column:action" json:"action"`
	MatchType           *string    `gorm:"column:match_type" json:"match_type"`
	ActionType          *string    `gorm:"column:action_type" json:"action_type"`
}

// TableName WebshellBlacklist's table name
func (*WebshellBlacklist) TableName() string {
	return TableNameWebshellBlacklist
}

func (*WebshellBlacklist) PrimaryKey() []string {
	return []string{}
}