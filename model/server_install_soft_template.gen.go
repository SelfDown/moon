// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameServerInstallSoftTemplate = "server_install_soft_template"

// ServerInstallSoftTemplate mapped from table <server_install_soft_template>
type ServerInstallSoftTemplate struct {
	SoftTemplateID           string     `gorm:"column:soft_template_id;primaryKey" json:"soft_template_id"`
	SoftTemplateName         string     `gorm:"column:soft_template_name;not null" json:"soft_template_name"`
	SoftTemplateType         string     `gorm:"column:soft_template_type;not null" json:"soft_template_type"`
	ServicePort              *string    `gorm:"column:service_port" json:"service_port"`
	ServerID                 string     `gorm:"column:server_id;not null" json:"server_id"`
	SoftName                 string     `gorm:"column:soft_name;not null" json:"soft_name"`
	SoftType                 string     `gorm:"column:soft_type;not null" json:"soft_type"`
	SoftHome                 string     `gorm:"column:soft_home;not null" json:"soft_home"`
	Portoffset               *string    `gorm:"column:portoffset" json:"portoffset"`
	CreateTime               *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	ModifyTime               *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"`
	Comments                 *string    `gorm:"column:comments" json:"comments"`
	DockerContainerName      *string    `gorm:"column:docker_container_name" json:"docker_container_name"`
	StartSoftwareCommand     *string    `gorm:"column:start_software_command" json:"start_software_command"`
	StopSoftwareCommand      *string    `gorm:"column:stop_software_command" json:"stop_software_command"`
	RestartSoftwareCommand   *string    `gorm:"column:restart_software_command" json:"restart_software_command"`
	ServerSoftLogPath        *string    `gorm:"column:server_soft_log_path" json:"server_soft_log_path"`
	ServerSoftLogsID         *string    `gorm:"column:server_soft_logs_id" json:"server_soft_logs_id"`
	ServerOsUserID           *int32     `gorm:"column:server_os_user_id" json:"server_os_user_id"`
	SoftID                   *int32     `gorm:"column:soft_id" json:"soft_id"`
	SoftwareLiveCheckCommand string     `gorm:"column:software_live_check_command;not null" json:"software_live_check_command"`
	IsMainSoft               *string    `gorm:"column:is_main_soft" json:"is_main_soft"`
	IsDocker                 *string    `gorm:"column:is_docker" json:"is_docker"`
	LastUseTime              *time.Time `gorm:"column:last_use_time;default:CURRENT_TIMESTAMP" json:"last_use_time"`
}

// TableName ServerInstallSoftTemplate's table name
func (*ServerInstallSoftTemplate) TableName() string {
	return TableNameServerInstallSoftTemplate
}

func (*ServerInstallSoftTemplate) PrimaryKey() []string {
	return []string{"soft_template_id"}
}