// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProjectArea = "project_area"

// ProjectArea mapped from table <project_area>
type ProjectArea struct {
	ProjectAreaID   string  `gorm:"column:project_area_id;primaryKey" json:"project_area_id"`
	Name            string  `gorm:"column:name;not null" json:"name"`
	ProjectCodeList string  `gorm:"column:project_code_list;not null" json:"project_code_list"`
	CreateUser      *string `gorm:"column:create_user" json:"create_user"`
	CreateTime      *string `gorm:"column:create_time" json:"create_time"`
}

// TableName ProjectArea's table name
func (*ProjectArea) TableName() string {
	return TableNameProjectArea
}

func (*ProjectArea) PrimaryKey() []string {
	return []string{"project_area_id"}
}