// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePublishWikiList = "publish_wiki_list"

// PublishWikiList mapped from table <publish_wiki_list>
type PublishWikiList struct {
	ID         int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ReqWikiID  string     `gorm:"column:req_wiki_id;not null" json:"req_wiki_id"`
	ReqID      string     `gorm:"column:req_id;not null" json:"req_id"`
	WikiKey    *string    `gorm:"column:wiki_key" json:"wiki_key"`
	WikiTitle  *string    `gorm:"column:wiki_title" json:"wiki_title"`
	WikiType   *int32     `gorm:"column:wiki_type" json:"wiki_type"` // 0 测试报告
	WikiStatu  *string    `gorm:"column:wiki_statu" json:"wiki_statu"`
	WikiURL    *string    `gorm:"column:wiki_url" json:"wiki_url"`                                          // 文档链接地址
	CreateTime *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments   *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName PublishWikiList's table name
func (*PublishWikiList) TableName() string {
	return TableNamePublishWikiList
}

func (*PublishWikiList) PrimaryKey() []string {
	return []string{"id"}
}