// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFailureReportChangeList = "failure_report_change_list"

// FailureReportChangeList mapped from table <failure_report_change_list>
type FailureReportChangeList struct {
	FailureReportChangeListID string  `gorm:"column:failure_report_change_list_id;primaryKey" json:"failure_report_change_list_id"`
	FailureReportID           *string `gorm:"column:failure_report_id" json:"failure_report_id"`
	Name                      string  `gorm:"column:name;not null" json:"name"`
	Operation                 *string `gorm:"column:operation" json:"operation"`
	Before                    *string `gorm:"column:before" json:"before"`
	After                     *string `gorm:"column:after" json:"after"`
	CreateUser                *string `gorm:"column:create_user" json:"create_user"`
	CreateTime                *string `gorm:"column:create_time" json:"create_time"`
}

// TableName FailureReportChangeList's table name
func (*FailureReportChangeList) TableName() string {
	return TableNameFailureReportChangeList
}

func (*FailureReportChangeList) PrimaryKey() []string {
	return []string{"failure_report_change_list_id"}
}