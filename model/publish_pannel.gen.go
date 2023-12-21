// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePublishPannel = "publish_pannel"

// PublishPannel mapped from table <publish_pannel>
type PublishPannel struct {
	PannelID          string  `gorm:"column:pannel_id;primaryKey" json:"pannel_id"`
	Type              *string `gorm:"column:type" json:"type"`
	Name              *string `gorm:"column:name" json:"name"`
	Spaces            *string `gorm:"column:spaces" json:"spaces"`
	IsLinkProductLine *string `gorm:"column:is_link_product_line" json:"is_link_product_line"`
}

// TableName PublishPannel's table name
func (*PublishPannel) TableName() string {
	return TableNamePublishPannel
}

func (*PublishPannel) PrimaryKey() []string {
	return []string{"pannel_id"}
}