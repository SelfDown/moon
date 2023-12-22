// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDocGroup = "doc_group"

// DocGroup mapped from table <doc_group>
type DocGroup struct {
	DocGroupID string  `gorm:"column:doc_group_id;primaryKey" json:"doc_group_id"`
	Name       *string `gorm:"column:name" json:"name"`
	Type       *string `gorm:"column:type" json:"type"`
	Desc       *string `gorm:"column:desc" json:"desc"`
	OrderIndex *int32  `gorm:"column:order_index" json:"order_index"`
	CreateTime *string `gorm:"column:create_time" json:"create_time"`
	CreateUser *string `gorm:"column:create_user" json:"create_user"`
	IsDelete   *string `gorm:"column:is_delete" json:"is_delete"`
}

// TableName DocGroup's table name
func (*DocGroup) TableName() string {
	return TableNameDocGroup
}

func (*DocGroup) PrimaryKey() []string {
	return []string{"doc_group_id"}
}