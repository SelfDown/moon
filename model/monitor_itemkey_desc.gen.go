// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMonitorItemkeyDesc = "monitor_itemkey_desc"

// MonitorItemkeyDesc mapped from table <monitor_itemkey_desc>
type MonitorItemkeyDesc struct {
	MonitorItemkeyDescID string  `gorm:"column:monitor_itemkey_desc_id;not null" json:"monitor_itemkey_desc_id"`
	Itemkey              *string `gorm:"column:itemkey" json:"itemkey"`
	ItemkeyDesc          *string `gorm:"column:itemkey_desc" json:"itemkey_desc"`
}

// TableName MonitorItemkeyDesc's table name
func (*MonitorItemkeyDesc) TableName() string {
	return TableNameMonitorItemkeyDesc
}

func (*MonitorItemkeyDesc) PrimaryKey() []string {
	return []string{}
}