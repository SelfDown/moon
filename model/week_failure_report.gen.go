// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWeekFailureReport = "week_failure_report"

// WeekFailureReport mapped from table <week_failure_report>
type WeekFailureReport struct {
	WeekFailureReportID string `gorm:"column:week_failure_report_id;primaryKey" json:"week_failure_report_id"`
	BeanReportID        string `gorm:"column:bean_report_id;not null" json:"bean_report_id"`
	FailureReportID     string `gorm:"column:failure_report_id;not null" json:"failure_report_id"`
}

// TableName WeekFailureReport's table name
func (*WeekFailureReport) TableName() string {
	return TableNameWeekFailureReport
}

func (*WeekFailureReport) PrimaryKey() []string {
	return []string{"week_failure_report_id"}
}