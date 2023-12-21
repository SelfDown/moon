// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDbdataImport = "dbdata_import"

// DbdataImport mapped from table <dbdata_import>
type DbdataImport struct {
	DbdataImportID      string  `gorm:"column:dbdata_import_id;primaryKey" json:"dbdata_import_id"`
	Title               *string `gorm:"column:title" json:"title"`
	SftpDumpPath        *string `gorm:"column:sftp_dump_path" json:"sftp_dump_path"`
	DbdataExportID      string  `gorm:"column:dbdata_export_id;not null" json:"dbdata_export_id"`
	InstallSoftID       *string `gorm:"column:install_soft_id" json:"install_soft_id"`
	SoftUserID          *string `gorm:"column:soft_user_id" json:"soft_user_id"`
	AddTime             *string `gorm:"column:add_time" json:"add_time"`
	OpUser              string  `gorm:"column:op_user;not null" json:"op_user"`
	Note                *string `gorm:"column:note" json:"note"`
	UploadFileName      *string `gorm:"column:upload_file_name" json:"upload_file_name"`
	UploadFilePath      *string `gorm:"column:upload_file_path" json:"upload_file_path"`
	ImportWay           *string `gorm:"column:import_way" json:"import_way"`       // 导出方式
	UserMetadata        *string `gorm:"column:user_metadata" json:"user_metadata"` // 用户表空间导出方式
	DirectoryPath       *string `gorm:"column:directory_path" json:"directory_path"`
	DirectoryName       *string `gorm:"column:directory_name" json:"directory_name"`
	OpPercent           *string `gorm:"column:op_percent" json:"op_percent"`
	CanEditor           *string `gorm:"column:can_editor" json:"can_editor"`
	StransactionIdx     *string `gorm:"column:stransaction_idx" json:"stransaction_idx"`
	ExpOraVersion       *string `gorm:"column:exp_ora_version" json:"exp_ora_version"`
	AllMetaDumpFilename *string `gorm:"column:all_meta_dump_filename" json:"all_meta_dump_filename"`
	Sid                 *string `gorm:"column:sid" json:"sid"`
	AllMetaLogFilename  *string `gorm:"column:all_meta_log_filename" json:"all_meta_log_filename"`
	IsDelete            *string `gorm:"column:is_delete;default:0" json:"is_delete"`
}

// TableName DbdataImport's table name
func (*DbdataImport) TableName() string {
	return TableNameDbdataImport
}

func (*DbdataImport) PrimaryKey() []string {
	return []string{"dbdata_import_id"}
}