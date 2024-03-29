// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameIssueImg = "issue_img"

// IssueImg mapped from table <issue_img>
type IssueImg struct {
	IssueImgID string  `gorm:"column:issue_img_id;primaryKey" json:"issue_img_id"`
	Data       *string `gorm:"column:data" json:"data"` // 文件内容
	AddTime    *string `gorm:"column:add_time" json:"add_time"`
	OpUser     string  `gorm:"column:op_user;not null" json:"op_user"`
}

// TableName IssueImg's table name
func (*IssueImg) TableName() string {
	return TableNameIssueImg
}

func (*IssueImg) PrimaryKey() []string {
	return []string{"issue_img_id"}
}