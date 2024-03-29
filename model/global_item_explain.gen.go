// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGlobalItemExplain = "global_item_explain"

// GlobalItemExplain mapped from table <global_item_explain>
type GlobalItemExplain struct {
	GlobalItemExplainID string     `gorm:"column:global_item_explain_id;primaryKey" json:"global_item_explain_id"`
	GlobalKey           *string    `gorm:"column:global_key" json:"global_key"`
	SysProjectID        *string    `gorm:"column:sys_project_id" json:"sys_project_id"` // 所应用的项目点区域
	ChangeEnvMod        *string    `gorm:"column:change_env_mod" json:"change_env_mod"` // 在其他环境使用是否需要调整
	FlagDel             *string    `gorm:"column:flag_del" json:"flag_del"`
	GlobalKeyDesc       *string    `gorm:"column:global_key_desc" json:"global_key_desc"`                            // key 的描述
	CreateTime          *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime          *time.Time `gorm:"column:modify_time" json:"modify_time"`
	OpUser              *string    `gorm:"column:op_user" json:"op_user"`
	Note                *string    `gorm:"column:note" json:"note"`
	Comments            *string    `gorm:"column:comments" json:"comments"` // 备注说明
}

// TableName GlobalItemExplain's table name
func (*GlobalItemExplain) TableName() string {
	return TableNameGlobalItemExplain
}

func (*GlobalItemExplain) PrimaryKey() []string {
	return []string{"global_item_explain_id"}
}