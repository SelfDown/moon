// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFrExerecord = "fr_exerecord"

// FrExerecord mapped from table <fr_exerecord>
type FrExerecord struct {
	ID       int32      `gorm:"column:id;not null" json:"id"`
	Tname    *string    `gorm:"column:tname" json:"tname"`
	Type     *int32     `gorm:"column:type" json:"type"`
	Param    *string    `gorm:"column:param" json:"param"`
	Logtime  *time.Time `gorm:"column:logtime" json:"logtime"`
	Project  *string    `gorm:"column:project" json:"project"`
	IP       *string    `gorm:"column:ip" json:"ip"`
	Username *string    `gorm:"column:username" json:"username"`
	Userrole *string    `gorm:"column:userrole" json:"userrole"`
	Time     *int32     `gorm:"column:time" json:"time"`
	Sql      *string    `gorm:"column:sql" json:"sql"`
	Browser  *string    `gorm:"column:browser" json:"browser"`
	Memory   *float64   `gorm:"column:memory" json:"memory"`
}

// TableName FrExerecord's table name
func (*FrExerecord) TableName() string {
	return TableNameFrExerecord
}

func (*FrExerecord) PrimaryKey() []string {
	return []string{}
}