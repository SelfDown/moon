// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSqlAdminLog = "sql_admin_log"

// SqlAdminLog mapped from table <sql_admin_log>
type SqlAdminLog struct {
	ID               int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID           *string `gorm:"column:user_id" json:"user_id"`
	SqlAdminLogID    string  `gorm:"column:sql_admin_log_id;not null" json:"sql_admin_log_id"` // 主键
	ServerEnvID      *string `gorm:"column:server_env_id" json:"server_env_id"`
	OracleInstanceID *string `gorm:"column:oracle_instance_id" json:"oracle_instance_id"` // 关联 oracle_instance.oracle_instance_id
	SqlAdminRoleID   *string `gorm:"column:sql_admin_role_id" json:"sql_admin_role_id"`   // 关联 sql_admin_role.sql_admin_role_id
	ExecSql          *string `gorm:"column:exec_sql" json:"exec_sql"`                     // 执行的SQL
	ExecTake         *string `gorm:"column:exec_take" json:"exec_take"`                   // 执行SQL消耗的时间 单位是秒
	/*
		执行结果
		成功
		失败
	*/
	ExecStatus    *string    `gorm:"column:exec_status" json:"exec_status"`
	ExecDetail    *string    `gorm:"column:exec_detail" json:"exec_detail"`                                    // 执行结果
	CreateTime    *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime    *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments      *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
	InstallSoftID *string    `gorm:"column:install_soft_id" json:"install_soft_id"`
}

// TableName SqlAdminLog's table name
func (*SqlAdminLog) TableName() string {
	return TableNameSqlAdminLog
}

func (*SqlAdminLog) PrimaryKey() []string {
	return []string{"id"}
}