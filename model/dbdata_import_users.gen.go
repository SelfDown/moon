// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDbdataImportUsers = "dbdata_import_users"

// DbdataImportUsers mapped from table <dbdata_import_users>
type DbdataImportUsers struct {
	DbdataImportUsersID string  `gorm:"column:dbdata_import_users_id;primaryKey" json:"dbdata_import_users_id"`
	DbdataImportID      string  `gorm:"column:dbdata_import_id;not null" json:"dbdata_import_id"`
	DbUser              *string `gorm:"column:db_user" json:"db_user"`
}

// TableName DbdataImportUsers's table name
func (*DbdataImportUsers) TableName() string {
	return TableNameDbdataImportUsers
}

func (*DbdataImportUsers) PrimaryKey() []string {
	return []string{"dbdata_import_users_id"}
}