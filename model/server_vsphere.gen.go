// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameServerVsphere = "server_vsphere"

// ServerVsphere mapped from table <server_vsphere>
type ServerVsphere struct {
	ServerVsphereID int32 `gorm:"column:server_vsphere_id;primaryKey;autoIncrement:true" json:"server_vsphere_id"` // 主键\\n
	/*
		环境ID

	*/
	ServerEnvID     *string `gorm:"column:server_env_id" json:"server_env_id"`
	VsphereServerIP *string `gorm:"column:vsphere_server_ip" json:"vsphere_server_ip"` // vsphere server ip
	VsphereUser     *string `gorm:"column:vsphere_user" json:"vsphere_user"`           // vsphere用户名和密码
	/*
		vsphere登陆密码

	*/
	VspherePwd        *string    `gorm:"column:vsphere_pwd" json:"vsphere_pwd"`
	DefaultDatacenter *string    `gorm:"column:default_datacenter" json:"default_datacenter"`                      // 默认数据中心
	DefaultCluster    *string    `gorm:"column:default_cluster" json:"default_cluster"`                            // 默认集群
	CreateTime        *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime        *time.Time `gorm:"column:modify_time" json:"modify_time"`
	Comment           *string    `gorm:"column:comment" json:"comment"`
	Comments          *string    `gorm:"column:comments" json:"comments"` // 备注说明
}

// TableName ServerVsphere's table name
func (*ServerVsphere) TableName() string {
	return TableNameServerVsphere
}

func (*ServerVsphere) PrimaryKey() []string {
	return []string{"server_vsphere_id"}
}