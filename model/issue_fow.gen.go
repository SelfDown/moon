// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameIssueFow = "issue_fow"

// IssueFow mapped from table <issue_fow>
type IssueFow struct {
	IssueFowID string `gorm:"column:issue_fow_id;primaryKey" json:"issue_fow_id"`
	IssueStatu string `gorm:"column:issue_statu;not null" json:"issue_statu"`
	Weight     *int32 `gorm:"column:weight" json:"weight"`
}

// TableName IssueFow's table name
func (*IssueFow) TableName() string {
	return TableNameIssueFow
}

func (*IssueFow) PrimaryKey() []string {
	return []string{"issue_fow_id"}
}
