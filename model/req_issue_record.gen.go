// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameReqIssueRecord = "req_issue_record"

// ReqIssueRecord mapped from table <req_issue_record>
type ReqIssueRecord struct {
	ReqIssueRecordID string  `gorm:"column:req_issue_record_id;primaryKey" json:"req_issue_record_id"`
	Issues           *string `gorm:"column:issues" json:"issues"` // 以逗号分割
	ServerEnvID      *string `gorm:"column:server_env_id" json:"server_env_id"`
	Description      string  `gorm:"column:description;not null" json:"description"` // 问题描述
	Comment          *string `gorm:"column:comment" json:"comment"`                  // 备注
	PublishVersion   *string `gorm:"column:publish_version" json:"publish_version"`  // 版本号
	SubmitUser       *string `gorm:"column:submit_user" json:"submit_user"`
	SubmitTime       *string `gorm:"column:submit_time" json:"submit_time"`
	IsBug            *string `gorm:"column:is_bug" json:"is_bug"`
	ProblemType      *string `gorm:"column:problem_type" json:"problem_type"`
	ResolveStatus    *string `gorm:"column:resolve_status" json:"resolve_status"`
	ResolveUser      *string `gorm:"column:resolve_user" json:"resolve_user"`
	ResolveTime      *string `gorm:"column:resolve_time" json:"resolve_time"`
	ResolveCount     *int32  `gorm:"column:resolve_count" json:"resolve_count"`
	AddTime          *string `gorm:"column:add_time" json:"add_time"`
	OpUser           string  `gorm:"column:op_user;not null" json:"op_user"`
	ProblemLevel     *string `gorm:"column:problem_level" json:"problem_level"`
	IsEffectVersion  *string `gorm:"column:is_effect_version" json:"is_effect_version"`
	Avoid            *string `gorm:"column:avoid" json:"avoid"`
	TestUser         *string `gorm:"column:test_user" json:"test_user"`
	RelativeIssues   *string `gorm:"column:relative_issues" json:"relative_issues"`
	RelativeVersion  *string `gorm:"column:relative_version" json:"relative_version"`
	RelativeReason   *string `gorm:"column:relative_reason" json:"relative_reason"`
	ProblemReasons   *string `gorm:"column:problem_reasons" json:"problem_reasons"`
}

// TableName ReqIssueRecord's table name
func (*ReqIssueRecord) TableName() string {
	return TableNameReqIssueRecord
}

func (*ReqIssueRecord) PrimaryKey() []string {
	return []string{"req_issue_record_id"}
}