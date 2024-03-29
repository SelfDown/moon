// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysProjectTeam = "sys_project_team"

// SysProjectTeam mapped from table <sys_project_team>
type SysProjectTeam struct {
	SysProjectTeamID string     `gorm:"column:sys_project_team_id;primaryKey" json:"sys_project_team_id"`
	TeamName         *string    `gorm:"column:team_name" json:"team_name"`
	TeamCode         *string    `gorm:"column:team_code" json:"team_code"`
	TeamLeader       *string    `gorm:"column:team_leader" json:"team_leader"`
	Notes            *string    `gorm:"column:notes" json:"notes"`
	TeamLeaderID     *string    `gorm:"column:team_leader_id" json:"team_leader_id"`                              // 产品负责人ID关联 sys_project_team_leader
	CreateTime       *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime       *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments         *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName SysProjectTeam's table name
func (*SysProjectTeam) TableName() string {
	return TableNameSysProjectTeam
}

func (*SysProjectTeam) PrimaryKey() []string {
	return []string{"sys_project_team_id"}
}