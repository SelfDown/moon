// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTmp = "tmp"

// Tmp mapped from table <tmp>
type Tmp struct {
	MonitorItemkeyDescID string  `gorm:"column:monitor_itemkey_desc_id;not null" json:"monitor_itemkey_desc_id"`
	Itemkey              *string `gorm:"column:itemkey" json:"itemkey"`
	ItemkeyDesc          *string `gorm:"column:itemkey_desc" json:"itemkey_desc"`
}

// TableName Tmp's table name
func (*Tmp) TableName() string {
	return TableNameTmp
}

func (*Tmp) PrimaryKey() []string {
	return []string{}
}