// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRecordCare = "record_care"

// RecordCare mapped from table <record_care>
type RecordCare struct {
	RecordCareID     string `gorm:"column:record_care_id;primaryKey" json:"record_care_id"`
	HisIssueRecordID string `gorm:"column:his_issue_record_id;not null" json:"his_issue_record_id"`
	WorkCode         string `gorm:"column:work_code;not null" json:"work_code"`
	CreateTime       string `gorm:"column:create_time;not null" json:"create_time"`
}

// TableName RecordCare's table name
func (*RecordCare) TableName() string {
	return TableNameRecordCare
}

func (*RecordCare) PrimaryKey() []string {
	return []string{"record_care_id"}
}