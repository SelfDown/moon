// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameIconTag = "icon_tag"

// IconTag mapped from table <icon_tag>
type IconTag struct {
	IconTagID  string  `gorm:"column:icon_tag_id;primaryKey" json:"icon_tag_id"`
	Tag        *string `gorm:"column:tag" json:"tag"`
	IconDataID *string `gorm:"column:icon_data_id" json:"icon_data_id"`
}

// TableName IconTag's table name
func (*IconTag) TableName() string {
	return TableNameIconTag
}

func (*IconTag) PrimaryKey() []string {
	return []string{"icon_tag_id"}
}