// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePlanConfirm = "plan_confirm"

// PlanConfirm mapped from table <plan_confirm>
type PlanConfirm struct {
	ReqConfirmID string  `gorm:"column:req_confirm_id;primaryKey" json:"req_confirm_id"`
	ServerEnvID  *string `gorm:"column:server_env_id" json:"server_env_id"`
	ProjectType  *string `gorm:"column:project_type" json:"project_type"`
	CreateTime   *string `gorm:"column:create_time" json:"create_time"`
	OpUser       *string `gorm:"column:op_user" json:"op_user"`
	ReqID        *string `gorm:"column:req_id" json:"req_id"`
}

// TableName PlanConfirm's table name
func (*PlanConfirm) TableName() string {
	return TableNamePlanConfirm
}

func (*PlanConfirm) PrimaryKey() []string {
	return []string{"req_confirm_id"}
}