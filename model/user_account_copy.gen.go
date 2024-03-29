// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserAccountCopy = "user_account_copy"

// UserAccountCopy mapped from table <user_account_copy>
type UserAccountCopy struct {
	Userid               string     `gorm:"column:userid;primaryKey" json:"userid"`
	RoleID               string     `gorm:"column:role_id;not null" json:"role_id"`
	Username             string     `gorm:"column:username;not null" json:"username"`
	Userpwd              string     `gorm:"column:userpwd;not null" json:"userpwd"`
	LastLoginTime        *time.Time `gorm:"column:last_login_time" json:"last_login_time"` // 最后一次登录时间
	LastLoginIP          *string    `gorm:"column:last_login_ip" json:"last_login_ip"`     // 最后一次登录IP
	Email                *string    `gorm:"column:email" json:"email"`                     // 用户邮箱
	Nick                 *string    `gorm:"column:nick" json:"nick"`                       // 用户昵称
	Statu                *int32     `gorm:"column:statu" json:"statu"`                     // 1 启用，0禁用
	Address              *string    `gorm:"column:address" json:"address"`                 // 物理地址
	Note                 *string    `gorm:"column:note" json:"note"`
	Email2               *string    `gorm:"column:email2" json:"email2"`                                              // 用户备用邮箱
	CreateTime           *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	LastLoginFailureTime *time.Time `gorm:"column:last_login_failure_time" json:"last_login_failure_time"`            // 最近一次登录失败时间
	LoginFailureCount    *int32     `gorm:"column:login_failure_count" json:"login_failure_count"`                    // 登录失败计数
	Avatar               *string    `gorm:"column:avatar" json:"avatar"`                                              // 用户头像base64
	ModifyTime           *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments             *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName UserAccountCopy's table name
func (*UserAccountCopy) TableName() string {
	return TableNameUserAccountCopy
}

func (*UserAccountCopy) PrimaryKey() []string {
	return []string{"userid"}
}