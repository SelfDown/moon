// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePublishIssueList = "publish_issue_list"

// PublishIssueList mapped from table <publish_issue_list>
type PublishIssueList struct {
	ID         int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ReqIssueID string  `gorm:"column:req_issue_id;not null" json:"req_issue_id"`
	ReqID      string  `gorm:"column:req_id;not null" json:"req_id"`
	IssueKey   *string `gorm:"column:issue_key" json:"issue_key"`
	IssueTitle *string `gorm:"column:issue_title" json:"issue_title"`
	IssueType  *int32  `gorm:"column:issue_type" json:"issue_type"` // 0 开发类，1，项目类，2-环境类
	IssueStatu *string `gorm:"column:issue_statu" json:"issue_statu"`
	Issuetype  *string `gorm:"column:issuetype" json:"issuetype"` // Jira的ISSUE问题类型
	/*
		是否为开始集成测试后补录的ISSUE

		0 或空 表示 非补录
		1 表示补录
	*/
	IsSupply      *string    `gorm:"column:isSupply" json:"isSupply"`
	IssueAssignee *string    `gorm:"column:issue_assignee" json:"issue_assignee"`                              // ISSUE经办人
	CreateTime    *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime    *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments      *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName PublishIssueList's table name
func (*PublishIssueList) TableName() string {
	return TableNamePublishIssueList
}

func (*PublishIssueList) PrimaryKey() []string {
	return []string{"id"}
}