// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameIconGroup = "icon_group"

// IconGroup mapped from table <icon_group>
type IconGroup struct {
	IconGroupID   string  `gorm:"column:icon_group_id;primaryKey" json:"icon_group_id"`
	IconProjectID *string `gorm:"column:icon_project_id" json:"icon_project_id"`
	GroupName     *string `gorm:"column:group_name" json:"group_name"`
	CreateTime    *string `gorm:"column:create_time" json:"create_time"`
}

// TableName IconGroup's table name
func (*IconGroup) TableName() string {
	return TableNameIconGroup
}

func (*IconGroup) PrimaryKey() []string {
	return []string{"icon_group_id"}
}