// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameIPIdentifyMiddle = "ip_identify_middle"

// IPIdentifyMiddle mapped from table <ip_identify_middle>
type IPIdentifyMiddle struct {
	IPIdentifyMiddleID int32   `gorm:"column:ip_identify_middle_id;primaryKey;autoIncrement:true" json:"ip_identify_middle_id"`
	BackupIP           string  `gorm:"column:backup_ip;not null" json:"backup_ip"`
	Identify           string  `gorm:"column:identify;not null" json:"identify"`
	SourceIP           *string `gorm:"column:source_ip" json:"source_ip"`
}

// TableName IPIdentifyMiddle's table name
func (*IPIdentifyMiddle) TableName() string {
	return TableNameIPIdentifyMiddle
}

func (*IPIdentifyMiddle) PrimaryKey() []string {
	return []string{"ip_identify_middle_id"}
}