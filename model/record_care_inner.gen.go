// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRecordCareInner = "record_care_inner"

// RecordCareInner mapped from table <record_care_inner>
type RecordCareInner struct {
	RecordCareInnerID string `gorm:"column:record_care_inner_id;primaryKey" json:"record_care_inner_id"`
	HisIssueRecordID  string `gorm:"column:his_issue_record_id;not null" json:"his_issue_record_id"`
	UserID            string `gorm:"column:user_id;not null" json:"user_id"`
	CreateTime        string `gorm:"column:create_time;not null" json:"create_time"`
}

// TableName RecordCareInner's table name
func (*RecordCareInner) TableName() string {
	return TableNameRecordCareInner
}

func (*RecordCareInner) PrimaryKey() []string {
	return []string{"record_care_inner_id"}
}