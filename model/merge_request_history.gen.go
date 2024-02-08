// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMergeRequestHistory = "merge_request_history"

// MergeRequestHistory mapped from table <merge_request_history>
type MergeRequestHistory struct {
	MergeRequestHistoryID string  `gorm:"column:merge_request_history_id;primaryKey" json:"merge_request_history_id"`
	GitlabURL             *string `gorm:"column:gitlab_url" json:"gitlab_url"`
	ProjectID             string  `gorm:"column:project_id;not null" json:"project_id"`
	FromBranch            string  `gorm:"column:from_branch;not null" json:"from_branch"`
	ToBranch              string  `gorm:"column:to_branch;not null" json:"to_branch"`
	MergeRequestID        *string `gorm:"column:merge_request_id" json:"merge_request_id"` // mr 的id
	Status                *string `gorm:"column:status" json:"status"`                     // 状态
	Message               *string `gorm:"column:message" json:"message"`                   // 日志
	CreateUser            *string `gorm:"column:create_user" json:"create_user"`           // 用户ID
	CreateTime            string  `gorm:"column:create_time;not null" json:"create_time"`
}

// TableName MergeRequestHistory's table name
func (*MergeRequestHistory) TableName() string {
	return TableNameMergeRequestHistory
}

func (*MergeRequestHistory) PrimaryKey() []string {
	return []string{"merge_request_history_id"}
}