// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMyFocusDashboard = "my_focus_dashboard"

// MyFocusDashboard mapped from table <my_focus_dashboard>
type MyFocusDashboard struct {
	MyDashboardID string     `gorm:"column:my_dashboard_id;primaryKey" json:"my_dashboard_id"` // ID
	DashboardID   *string    `gorm:"column:dashboard_id" json:"dashboard_id"`                  // 面板ID
	AddTime       *time.Time `gorm:"column:add_time" json:"add_time"`
	Userid        *string    `gorm:"column:userid" json:"userid"` // 用户主键ID
}

// TableName MyFocusDashboard's table name
func (*MyFocusDashboard) TableName() string {
	return TableNameMyFocusDashboard
}

func (*MyFocusDashboard) PrimaryKey() []string {
	return []string{"my_dashboard_id"}
}