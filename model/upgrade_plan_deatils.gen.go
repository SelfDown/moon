// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUpgradePlanDeatils = "upgrade_plan_deatils"

// UpgradePlanDeatils mapped from table <upgrade_plan_deatils>
type UpgradePlanDeatils struct {
	UpgradePlanDetilsID string     `gorm:"column:upgrade_plan_detils_id;primaryKey" json:"upgrade_plan_detils_id"`   // UUID--用他做主键
	UpgradePlanListID   *string    `gorm:"column:upgrade_plan_list_id" json:"upgrade_plan_list_id"`                  // 主表ID
	ArtifactSource      *string    `gorm:"column:artifact_source" json:"artifact_source"`                            // 需要升级的war包名称
	VersionSource       *string    `gorm:"column:version_source" json:"version_source"`                              // 需要升级的war包版本
	ArtifactTarget      *string    `gorm:"column:artifact_target" json:"artifact_target"`                            // 对比目标war包名称
	VersionTarget       *string    `gorm:"column:version_target" json:"version_target"`                              // 对比目标(参考对象)war包版本
	IsRight             *string    `gorm:"column:isRight" json:"isRight"`                                            // 0- 表示这条记录用户不关注（下载关联脚本的时候不会下载） 1表示用户关注，对于他有用
	ReqID               *string    `gorm:"column:req_id" json:"req_id"`                                              // 关联升级单主表
	ReleaseTime         *time.Time `gorm:"column:release_time" json:"release_time"`                                  // war包升级发布时间
	CreateTime          *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime          *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments            *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName UpgradePlanDeatils's table name
func (*UpgradePlanDeatils) TableName() string {
	return TableNameUpgradePlanDeatils
}

func (*UpgradePlanDeatils) PrimaryKey() []string {
	return []string{"upgrade_plan_detils_id"}
}