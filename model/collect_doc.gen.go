// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCollectDoc = "collect_doc"

// CollectDoc mapped from table <collect_doc>
type CollectDoc struct {
	CollectDocID string  `gorm:"column:collect_doc_id;primaryKey" json:"collect_doc_id"`
	Title        *string `gorm:"column:title" json:"title"`
	SubTitle     *string `gorm:"column:sub_title" json:"sub_title"`
	Type         *string `gorm:"column:type" json:"type"`
	ParentDir    *string `gorm:"column:parent_dir" json:"parent_dir"`
	Code         *string `gorm:"column:code" json:"code"`
	CodeDesc     *string `gorm:"column:code_desc" json:"code_desc"`
	OrderIndex   *int32  `gorm:"column:order_index" json:"order_index"`
	CreateTime   *string `gorm:"column:create_time" json:"create_time"`
	CreateUser   *string `gorm:"column:create_user" json:"create_user"`
	IsDelete     *string `gorm:"column:is_delete" json:"is_delete"`
	CodeResult   *string `gorm:"column:code_result" json:"code_result"`
}

// TableName CollectDoc's table name
func (*CollectDoc) TableName() string {
	return TableNameCollectDoc
}

func (*CollectDoc) PrimaryKey() []string {
	return []string{"collect_doc_id"}
}