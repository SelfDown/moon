// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUsers = "users"

// Users mapped from table <users>
type Users struct {
	ID       int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username *string `gorm:"column:username" json:"username"`
	Password *string `gorm:"column:password" json:"password"`
}

// TableName Users's table name
func (*Users) TableName() string {
	return TableNameUsers
}

func (*Users) PrimaryKey() []string {
	return []string{"id"}
}