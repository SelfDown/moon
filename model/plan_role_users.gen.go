// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePlanRoleUsers = "plan_role_users"

// PlanRoleUsers mapped from table <plan_role_users>
type PlanRoleUsers struct {
	PlanRoleUsersID string  `gorm:"column:plan_role_users_id;primaryKey" json:"plan_role_users_id"`
	PlanID          string  `gorm:"column:plan_id;not null" json:"plan_id"`
	UserID          *string `gorm:"column:user_id" json:"user_id"` // 用户ID
	Type            string  `gorm:"column:type;not null" json:"type"`
}

// TableName PlanRoleUsers's table name
func (*PlanRoleUsers) TableName() string {
	return TableNamePlanRoleUsers
}

func (*PlanRoleUsers) PrimaryKey() []string {
	return []string{"plan_role_users_id"}
}