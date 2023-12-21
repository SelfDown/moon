// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameServerSoftConf = "server_soft_conf"

// ServerSoftConf mapped from table <server_soft_conf>
type ServerSoftConf struct {
	ID                  int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SoftName            *string `gorm:"column:soft_name" json:"soft_name"`
	DefaultSoftVersions *string `gorm:"column:default_soft_versions" json:"default_soft_versions"`
	SoftGitPath         *string `gorm:"column:soft_git_path" json:"soft_git_path"`
	SoftHomePath        *string `gorm:"column:soft_home_path" json:"soft_home_path"`
}

// TableName ServerSoftConf's table name
func (*ServerSoftConf) TableName() string {
	return TableNameServerSoftConf
}

func (*ServerSoftConf) PrimaryKey() []string {
	return []string{"id"}
}