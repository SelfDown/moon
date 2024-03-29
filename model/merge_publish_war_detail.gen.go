// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMergePublishWarDetail = "merge_publish_war_detail"

// MergePublishWarDetail mapped from table <merge_publish_war_detail>
type MergePublishWarDetail struct {
	MergePublishWarDetailID string  `gorm:"column:merge_publish_war_detail_id;primaryKey" json:"merge_publish_war_detail_id"`
	MergePublishID          string  `gorm:"column:merge_publish_id;not null" json:"merge_publish_id"`
	PublishVersion          string  `gorm:"column:publish_version;not null" json:"publish_version"`
	WarArtifactid           string  `gorm:"column:war_artifactid;not null" json:"war_artifactid"`
	WarGroupid              *string `gorm:"column:war_groupid" json:"war_groupid"`
	WarVersion              *string `gorm:"column:war_version" json:"war_version"`
	CreateUser              *string `gorm:"column:create_user" json:"create_user"`
	CreateTime              *string `gorm:"column:create_time" json:"create_time"`
}

// TableName MergePublishWarDetail's table name
func (*MergePublishWarDetail) TableName() string {
	return TableNameMergePublishWarDetail
}

func (*MergePublishWarDetail) PrimaryKey() []string {
	return []string{"merge_publish_war_detail_id"}
}
