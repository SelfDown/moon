// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFrErrrecord = "fr_errrecord"

// FrErrrecord mapped from table <fr_errrecord>
type FrErrrecord struct {
	ID      int32      `gorm:"column:id;not null" json:"id"`
	Tname   *string    `gorm:"column:tname" json:"tname"`
	Sinfo   *string    `gorm:"column:sinfo" json:"sinfo"`
	Type    *int32     `gorm:"column:type" json:"type"`
	Msg     *string    `gorm:"column:msg" json:"msg"`
	Logtime *time.Time `gorm:"column:logtime" json:"logtime"`
	Project *string    `gorm:"column:project" json:"project"`
}

// TableName FrErrrecord's table name
func (*FrErrrecord) TableName() string {
	return TableNameFrErrrecord
}

func (*FrErrrecord) PrimaryKey() []string {
	return []string{}
}