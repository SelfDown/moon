// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDbdataExportUsers = "dbdata_export_users"

// DbdataExportUsers mapped from table <dbdata_export_users>
type DbdataExportUsers struct {
	DbdataExportUsersID string  `gorm:"column:dbdata_export_users_id;primaryKey" json:"dbdata_export_users_id"`
	DbdataExportID      string  `gorm:"column:dbdata_export_id;not null" json:"dbdata_export_id"`
	DbUser              *string `gorm:"column:db_user" json:"db_user"`
}

// TableName DbdataExportUsers's table name
func (*DbdataExportUsers) TableName() string {
	return TableNameDbdataExportUsers
}

func (*DbdataExportUsers) PrimaryKey() []string {
	return []string{"dbdata_export_users_id"}
}