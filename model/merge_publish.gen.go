// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMergePublish = "merge_publish"

// MergePublish mapped from table <merge_publish>
type MergePublish struct {
	MergePublishID string  `gorm:"column:merge_publish_id;primaryKey" json:"merge_publish_id"`
	MergeVersion   string  `gorm:"column:merge_version;not null" json:"merge_version"`
	MergeName      string  `gorm:"column:merge_name;not null" json:"merge_name"`
	PublishStatus  *string `gorm:"column:publish_status" json:"publish_status"`
	Comments       *string `gorm:"column:comments" json:"comments"`
	CreateUser     *string `gorm:"column:create_user" json:"create_user"`
	CreateTime     *string `gorm:"column:create_time" json:"create_time"`
	IsDelete       *string `gorm:"column:is_delete" json:"is_delete"`
}

// TableName MergePublish's table name
func (*MergePublish) TableName() string {
	return TableNameMergePublish
}

func (*MergePublish) PrimaryKey() []string {
	return []string{"merge_publish_id"}
}
