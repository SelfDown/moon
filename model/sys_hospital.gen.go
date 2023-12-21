// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysHospital = "sys_hospital"

// SysHospital mapped from table <sys_hospital>
type SysHospital struct {
	HospitalName *string    `gorm:"column:hospital_name" json:"hospital_name"`                                // 参数名称
	HospitalCode string     `gorm:"column:hospital_code;primaryKey" json:"hospital_code"`                     // 参数值
	GitBranchSql *string    `gorm:"column:git_branch_sql" json:"git_branch_sql"`                              // 院区对应的sql归档,git分支
	CreateTime   *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime   *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments     *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
	SysProjectID *string    `gorm:"column:sys_project_id" json:"sys_project_id"`
	OrderID      *int32     `gorm:"column:order_id" json:"order_id"`
}

// TableName SysHospital's table name
func (*SysHospital) TableName() string {
	return TableNameSysHospital
}

func (*SysHospital) PrimaryKey() []string {
	return []string{"hospital_code"}
}