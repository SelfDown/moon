// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePageGroup = "page_group"

// PageGroup mapped from table <page_group>
type PageGroup struct {
	PageGroupID    string `gorm:"column:page_group_id;primaryKey" json:"page_group_id"`
	Name           string `gorm:"column:name;not null" json:"name"` // 名称
	SysProjectCode string `gorm:"column:sys_project_code;not null" json:"sys_project_code"`
}

// TableName PageGroup's table name
func (*PageGroup) TableName() string {
	return TableNamePageGroup
}

func (*PageGroup) PrimaryKey() []string {
	return []string{"page_group_id"}
}