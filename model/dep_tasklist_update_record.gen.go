// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDepTasklistUpdateRecord = "dep_tasklist_update_record"

// DepTasklistUpdateRecord mapped from table <dep_tasklist_update_record>
type DepTasklistUpdateRecord struct {
	UpdateRecordID string  `gorm:"column:update_record_id;primaryKey" json:"update_record_id"`
	ServerEnvID    *string `gorm:"column:server_env_id" json:"server_env_id"`
	ServerID       *string `gorm:"column:server_id" json:"server_id"` // 服务器ID
	ServerIP       *string `gorm:"column:server_ip" json:"server_ip"` // 服务器IP
	WarGroupid     *string `gorm:"column:war_groupid" json:"war_groupid"`
	WarArtifactid  *string `gorm:"column:war_artifactid" json:"war_artifactid"`
	Field          *string `gorm:"column:field" json:"field"`
	Name           *string `gorm:"column:name" json:"name"`
	Before         *string `gorm:"column:before" json:"before"`
	After          *string `gorm:"column:after" json:"after"`
	CreateTime     *string `gorm:"column:create_time" json:"create_time"`
	OpUser         *string `gorm:"column:op_user" json:"op_user"`
	Operation      *string `gorm:"column:operation" json:"operation"`
}

// TableName DepTasklistUpdateRecord's table name
func (*DepTasklistUpdateRecord) TableName() string {
	return TableNameDepTasklistUpdateRecord
}

func (*DepTasklistUpdateRecord) PrimaryKey() []string {
	return []string{"update_record_id"}
}