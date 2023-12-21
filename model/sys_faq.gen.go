// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysFaq = "sys_faq"

// SysFaq mapped from table <sys_faq>
type SysFaq struct {
	SysFaqID    string     `gorm:"column:sys_faq_id;primaryKey" json:"sys_faq_id"`
	Question    *string    `gorm:"column:question" json:"question"`                                          // 问题
	Answer      *string    `gorm:"column:answer" json:"answer"`                                              // 问题答案
	CreateTime  *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime  *time.Time `gorm:"column:modify_time" json:"modify_time"`
	Comments    *string    `gorm:"column:comments" json:"comments"`
	AgreeCount  *int32     `gorm:"column:agree_count" json:"agree_count"`   // 问题被赞同的数量
	OpposeCount *int32     `gorm:"column:oppose_count" json:"oppose_count"` // 问题被不赞同的数量
}

// TableName SysFaq's table name
func (*SysFaq) TableName() string {
	return TableNameSysFaq
}

func (*SysFaq) PrimaryKey() []string {
	return []string{"sys_faq_id"}
}