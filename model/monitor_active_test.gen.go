// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMonitorActiveTest = "monitor_active_test"

// MonitorActiveTest mapped from table <monitor_active_test>
type MonitorActiveTest struct {
	ActiveTestID      string     `gorm:"column:active_test_id;not null" json:"active_test_id"`
	HostID            *string    `gorm:"column:host_id" json:"host_id"`
	HostName          *string    `gorm:"column:host_name" json:"host_name"`
	ActiveTestName    *string    `gorm:"column:active_test_name" json:"active_test_name"`
	TestIP            *string    `gorm:"column:test_ip" json:"test_ip"`
	TestPort          *string    `gorm:"column:test_port" json:"test_port"`
	TestURL           *string    `gorm:"column:test_url" json:"test_url"`
	TestSql           *string    `gorm:"column:test_sql" json:"test_sql"`
	TestType          string     `gorm:"column:test_type;not null" json:"test_type"`
	AddTime           *time.Time `gorm:"column:add_time;not null;default:CURRENT_TIMESTAMP" json:"add_time"`
	Applications      *string    `gorm:"column:applications" json:"applications"`
	Delay             *string    `gorm:"column:delay" json:"delay"`
	ValueType         *string    `gorm:"column:value_type" json:"value_type"`
	TriggerID         *string    `gorm:"column:trigger_id" json:"trigger_id"`
	RequestMethod     *string    `gorm:"column:request_method" json:"request_method"`
	TriggerName       *string    `gorm:"column:trigger_name" json:"trigger_name"`
	Key               *string    `gorm:"column:key_" json:"key_"`
	OutputFormat      *string    `gorm:"column:output_format" json:"output_format"`
	ItemID            *string    `gorm:"column:item_id" json:"item_id"`
	TestInstallSoftID *string    `gorm:"column:test_install_soft_id" json:"test_install_soft_id"`
	TestProxyParam    *string    `gorm:"column:test_proxy_param" json:"test_proxy_param"`
	TestSqlCode       *string    `gorm:"column:test_sql_code" json:"test_sql_code"`
	TestSqlDbuser     *string    `gorm:"column:test_sql_dbuser" json:"test_sql_dbuser"`
	TestSqlDbpwd      *string    `gorm:"column:test_sql_dbpwd" json:"test_sql_dbpwd"`
	SoftUserID        *string    `gorm:"column:soft_user_id" json:"soft_user_id"`
	ProjectCode       *string    `gorm:"column:project_code" json:"project_code"`
	ActiveTestTarget  *string    `gorm:"column:active_test_target" json:"active_test_target"`
	AlertitemCode     *string    `gorm:"column:alertitem_code" json:"alertitem_code"`
}

// TableName MonitorActiveTest's table name
func (*MonitorActiveTest) TableName() string {
	return TableNameMonitorActiveTest
}

func (*MonitorActiveTest) PrimaryKey() []string {
	return []string{}
}