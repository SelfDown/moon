// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameReqDeployTag = "req_deploy_tag"

// ReqDeployTag mapped from table <req_deploy_tag>
type ReqDeployTag struct {
	ReqDeployTagID string  `gorm:"column:req_deploy_tag_id;primaryKey" json:"req_deploy_tag_id"`
	ReqID          string  `gorm:"column:req_id;not null" json:"req_id"`
	SysProjectID   *string `gorm:"column:sys_project_id" json:"sys_project_id"` // 用户ID
}

// TableName ReqDeployTag's table name
func (*ReqDeployTag) TableName() string {
	return TableNameReqDeployTag
}

func (*ReqDeployTag) PrimaryKey() []string {
	return []string{"req_deploy_tag_id"}
}