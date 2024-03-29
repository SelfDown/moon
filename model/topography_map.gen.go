// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTopographyMap = "topography_map"

// TopographyMap mapped from table <topography_map>
type TopographyMap struct {
	TopographyMapID string `gorm:"column:topography_map_id;primaryKey" json:"topography_map_id"`
	ServerID        string `gorm:"column:server_id;not null" json:"server_id"`
	NodeType        string `gorm:"column:node_type;not null" json:"node_type"`
	Node            string `gorm:"column:node;not null" json:"node"`
}

// TableName TopographyMap's table name
func (*TopographyMap) TableName() string {
	return TableNameTopographyMap
}

func (*TopographyMap) PrimaryKey() []string {
	return []string{"topography_map_id"}
}