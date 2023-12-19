// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOpsConfigDetail = "ops_config_detail"

// OpsConfigDetail mapped from table <ops_config_detail>
type OpsConfigDetail struct {
	OpsConfigDetailID string     `gorm:"column:ops_config_detail_id;primaryKey" json:"ops_config_detail_id"`       // 主键
	OpsConfigKey      *string    `gorm:"column:ops_config_key" json:"ops_config_key"`                              // 键值对键
	OpsConfigValue    *string    `gorm:"column:ops_config_value" json:"ops_config_value"`                          // 键值对值
	OpsConfigGroupID  *string    `gorm:"column:ops_config_group_id" json:"ops_config_group_id"`                    // 所属配置分组 ops_config_group.ops_config_group_id
	CreateTime        *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime        *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments          *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
	/*
		软件配置文件相对路径
		比如： jboss/standalone/configruation/standalone.xml
	*/
	OpsConfigFileRelPath *string `gorm:"column:ops_config_file_rel_path" json:"ops_config_file_rel_path"`
}

// TableName OpsConfigDetail's table name
func (*OpsConfigDetail) TableName() string {
	return TableNameOpsConfigDetail
}

func (*OpsConfigDetail) PrimaryKey() []string {
	return []string{"ops_config_detail_id"}
}