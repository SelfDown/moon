// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID    int64   `gorm:"column:id;primaryKey" json:"id"` // 主键ID
	Name  *string `gorm:"column:name" json:"name"`        // 姓名
	Age   *int32  `gorm:"column:age" json:"age"`          // 年龄
	Email *string `gorm:"column:email" json:"email"`      // 邮箱
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

func (*User) PrimaryKey() []string {
	return []string{"id"}
}