// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDbdataImportTables = "dbdata_import_tables"

// DbdataImportTables mapped from table <dbdata_import_tables>
type DbdataImportTables struct {
	DbdataImportTablesID string  `gorm:"column:dbdata_import_tables_id;primaryKey" json:"dbdata_import_tables_id"`
	DbdataImportID       string  `gorm:"column:dbdata_import_id;not null" json:"dbdata_import_id"`
	DbUser               *string `gorm:"column:db_user" json:"db_user"`
	DbTable              *string `gorm:"column:db_table" json:"db_table"`
	DbQuery              *string `gorm:"column:db_query" json:"db_query"`
	ImportWay            *string `gorm:"column:import_way" json:"import_way"` // 导出方式
}

// TableName DbdataImportTables's table name
func (*DbdataImportTables) TableName() string {
	return TableNameDbdataImportTables
}

func (*DbdataImportTables) PrimaryKey() []string {
	return []string{"dbdata_import_tables_id"}
}