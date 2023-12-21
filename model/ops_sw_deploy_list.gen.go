// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOpsSwDeployList = "ops_sw_deploy_list"

// OpsSwDeployList mapped from table <ops_sw_deploy_list>
type OpsSwDeployList struct {
	OpsSwDeployListID string  `gorm:"column:ops_sw_deploy_list_id;primaryKey" json:"ops_sw_deploy_list_id"` // 主键
	ServerID          *string `gorm:"column:server_id" json:"server_id"`                                    // server_instance.server_id
	/*
		ops_config_group.ops_config_group_id
		软件配置分组
	*/
	OpsConfigGroupID   *string    `gorm:"column:ops_config_group_id" json:"ops_config_group_id"`
	OpsSoftwareListID  *string    `gorm:"column:ops_software_list_id" json:"ops_software_list_id"`
	CreateTime         *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime         *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments           *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
	OpsSoftwareVersion string     `gorm:"column:ops_software_version;not null" json:"ops_software_version"`         // 软件版本号
	InstallSoftID      *string    `gorm:"column:install_soft_id" json:"install_soft_id"`
}

// TableName OpsSwDeployList's table name
func (*OpsSwDeployList) TableName() string {
	return TableNameOpsSwDeployList
}

func (*OpsSwDeployList) PrimaryKey() []string {
	return []string{"ops_sw_deploy_list_id"}
}