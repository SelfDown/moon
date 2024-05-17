// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model



const TableNameServerEnv = "server_env"

// ServerEnv mapped from table <server_env>
type ServerEnv struct {
	ServerEnvID   string  `gorm:"column:server_env_id;primaryKey" json:"server_env_id"` // 服务器分组ID
	ServerEnvName *string `gorm:"column:server_env_name" json:"server_env_name"`        // 服务器分组名称
	Notes         *string `gorm:"column:notes" json:"notes"`                            // 备注
	HospitalCode  *string `gorm:"column:hospital_code" json:"hospital_code"`
	EnvCode       *string `gorm:"column:env_code" json:"env_code"`
	FlagDel       *string `gorm:"column:flag_del" json:"flag_del"`               // 逻辑删除 1删除 0 未删除
	Leader        *string `gorm:"column:leader" json:"leader"`                   // 环境负责人
	SysProjectID  *string `gorm:"column:sys_project_id" json:"sys_project_id"`   // 所属项目ID
	RoleAuthLevel *int32  `gorm:"column:role_auth_level" json:"role_auth_level"` // 角色权限等级
	EnvDocURL     *string `gorm:"column:env_doc_url" json:"env_doc_url"`         // 环境实例文档URL
	/*
		是否开启不间断部署
		0：禁用
		1：启用
	*/
	EnableHaDeploy *string    `gorm:"column:enable_ha_deploy" json:"enable_ha_deploy"`
	CreateTime     *string `gorm:"column:create_time;" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime     *string `gorm:"column:modify_time;" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments       *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
	LeaderTel      *string    `gorm:"column:leader_tel" json:"leader_tel"`
	OrderID        *int32     `gorm:"column:order_id" json:"order_id"`
	ServerEnvCode  *string    `gorm:"column:server_env_code" json:"server_env_code"`
	ServerGroupID  *string    `gorm:"column:server_group_id" json:"server_group_id"`
	ParentID       *string    `gorm:"column:parent_id" json:"parent_id"`
}

// TableName ServerEnv's table name
func (*ServerEnv) TableName() string {
	return TableNameServerEnv
}

func (*ServerEnv) PrimaryKey() []string {
	return []string{"server_env_id"}
}