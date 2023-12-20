// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCollectDoc = "collect_doc"

// CollectDoc mapped from table <collect_doc>
type CollectDoc struct {
	CollectDocID *string `gorm:"column:collect_doc_id;type:varchar(50)" json:"collect_doc_id"`
	Title        *string `gorm:"column:title;type:varchar(200)" json:"title"`
	SubTitle     *string `gorm:"column:sub_title;type:varchar(200)" json:"sub_title"`
	Type         *string `gorm:"column:type;type:varchar(200)" json:"type"`
	ParentDir    *string `gorm:"column:parent_dir;type:varchar(200)" json:"parent_dir"`
	Code         *string `gorm:"column:code;type:varchar(2550)" json:"code"`
	CodeDesc     *string `gorm:"column:code_desc;type:varchar(2550)" json:"code_desc"`
	OrderIndex   *string `gorm:"column:order_index;type:int(11)" json:"order_index"`
	CreateTime   *string `gorm:"column:create_time;type:varchar(50)" json:"create_time"`
	CreateUser   *string `gorm:"column:create_user;type:varchar(50)" json:"create_user"`
	IsDelete     *string `gorm:"column:is_delete;type:varchar(50)" json:"is_delete"`
	CodeResult   *string `gorm:"column:code_result;type:text" json:"code_result"`
}

// TableName CollectDoc's table name
func (*CollectDoc) TableName() string {
	return TableNameCollectDoc
}

func (*CollectDoc) PrimaryKey() []string {
	return []string{}
}