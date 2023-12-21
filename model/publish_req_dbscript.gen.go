// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePublishReqDbscript = "publish_req_dbscript"

// PublishReqDbscript mapped from table <publish_req_dbscript>
type PublishReqDbscript struct {
	ReqDbscriptID   string  `gorm:"column:req_dbscript_id;primaryKey" json:"req_dbscript_id"`
	ReqID           string  `gorm:"column:req_id;not null" json:"req_id"`
	DbscriptEventID *string `gorm:"column:dbscript_event_id" json:"dbscript_event_id"`
	/*
		是否为开始集成测试后补录的脚本归档

		0 或空 表示 非补录
		1 表示补录
	*/
	IsSupply    *string    `gorm:"column:isSupply" json:"isSupply"`
	CreateTime  *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime  *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments    *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
	SqlScriptID *string    `gorm:"column:sql_script_id" json:"sql_script_id"`
	OrderIndex  *string    `gorm:"column:order_index" json:"order_index"` // 执行排序字段
}

// TableName PublishReqDbscript's table name
func (*PublishReqDbscript) TableName() string {
	return TableNamePublishReqDbscript
}

func (*PublishReqDbscript) PrimaryKey() []string {
	return []string{"req_dbscript_id"}
}