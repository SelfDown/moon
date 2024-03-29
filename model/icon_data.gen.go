// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameIconData = "icon_data"

// IconData mapped from table <icon_data>
type IconData struct {
	IconDataID    string  `gorm:"column:icon_data_id;primaryKey" json:"icon_data_id"`
	IconGroupID   string  `gorm:"column:icon_group_id;not null" json:"icon_group_id"`
	IconProjectID *string `gorm:"column:icon_project_id" json:"icon_project_id"`
	DataType      *string `gorm:"column:data_type" json:"data_type"`
	Name          *string `gorm:"column:name" json:"name"`
	Pinyin        *string `gorm:"column:pinyin" json:"pinyin"`
	Data          string  `gorm:"column:data;not null" json:"data"`
	Status        *string `gorm:"column:status" json:"status"`
	CreateTime    *string `gorm:"column:create_time" json:"create_time"`
	CreateUser    *string `gorm:"column:create_user" json:"create_user"`
	IsDelete      *string `gorm:"column:is_delete" json:"is_delete"`
}

// TableName IconData's table name
func (*IconData) TableName() string {
	return TableNameIconData
}

func (*IconData) PrimaryKey() []string {
	return []string{"icon_data_id"}
}