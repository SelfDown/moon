// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDeployConfDiffDetail = "deploy_conf_diff_detail"

// DeployConfDiffDetail mapped from table <deploy_conf_diff_detail>
type DeployConfDiffDetail struct {
	DiffID        string     `gorm:"column:diff_id;primaryKey" json:"diff_id"`
	RepWarID      *string    `gorm:"column:rep_war_id" json:"rep_war_id"` // global参数对比临时表
	ParamKey      *string    `gorm:"column:param_key" json:"param_key"`
	ParamOrgValue *string    `gorm:"column:param_org_value" json:"param_org_value"`                            // 替换原始值
	ParamRepValue *string    `gorm:"column:param_rep_value" json:"param_rep_value"`                            // 替换后的值
	NotReplace    *int32     `gorm:"column:not_replace" json:"not_replace"`                                    // 是否被替换 1、被替换 0、未被替换 2 新增KEY
	CreateTime    *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime    *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments      *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName DeployConfDiffDetail's table name
func (*DeployConfDiffDetail) TableName() string {
	return TableNameDeployConfDiffDetail
}

func (*DeployConfDiffDetail) PrimaryKey() []string {
	return []string{"diff_id"}
}