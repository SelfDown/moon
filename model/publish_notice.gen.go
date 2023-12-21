// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePublishNotice = "publish_notice"

// PublishNotice mapped from table <publish_notice>
type PublishNotice struct {
	PublishNoticeID   string     `gorm:"column:publish_notice_id;primaryKey" json:"publish_notice_id"`
	PublishTime       *time.Time `gorm:"column:publish_time" json:"publish_time"`                                  // 申请单发布时间（小时：分）
	NoticeContent     *string    `gorm:"column:notice_content" json:"notice_content"`                              // 通知主体内容
	PublishUsetime    *string    `gorm:"column:publish_usetime" json:"publish_usetime"`                            // 预计用时（分）
	NoticeCreattime   *time.Time `gorm:"column:notice_creattime" json:"notice_creattime"`                          // 通知创建时间
	NoticeUserid      *string    `gorm:"column:notice_userid" json:"notice_userid"`                                // 创建者（存userid）
	MessageSendLogID  *string    `gorm:"column:message_send_log_id" json:"message_send_log_id"`                    // 发送消息ID
	ReqID             *string    `gorm:"column:req_id" json:"req_id"`                                              // 升级单ID
	PublishNoticeType *string    `gorm:"column:publish_notice_type" json:"publish_notice_type"`                    // 通知类型
	CreateTime        *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime        *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments          *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName PublishNotice's table name
func (*PublishNotice) TableName() string {
	return TableNamePublishNotice
}

func (*PublishNotice) PrimaryKey() []string {
	return []string{"publish_notice_id"}
}