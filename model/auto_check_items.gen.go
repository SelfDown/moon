// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAutoCheckItems = "auto_check_items"

// AutoCheckItems mapped from table <auto_check_items>
type AutoCheckItems struct {
	ItemsID       string     `gorm:"column:items_id;primaryKey" json:"items_id"`   // 监控分类的id，主键
	ItemsName     string     `gorm:"column:items_name;not null" json:"items_name"` // 监控分类的名称
	ItemsComments *string    `gorm:"column:items_comments" json:"items_comments"`
	CreateTime    *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime    *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments      *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName AutoCheckItems's table name
func (*AutoCheckItems) TableName() string {
	return TableNameAutoCheckItems
}

func (*AutoCheckItems) PrimaryKey() []string {
	return []string{"items_id"}
}