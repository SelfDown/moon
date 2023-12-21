// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameBeanReport = "bean_report"

// BeanReport mapped from table <bean_report>
type BeanReport struct {
	BeanReportID string  `gorm:"column:bean_report_id;primaryKey" json:"bean_report_id"`
	Name         *string `gorm:"column:name" json:"name"`      // 标题
	Yw           string  `gorm:"column:yw;not null" json:"yw"` // 按照年和周生成一个版本
	StartTime    *string `gorm:"column:start_time" json:"start_time"`
	EndTime      *string `gorm:"column:end_time" json:"end_time"`
	CreateTime   *string `gorm:"column:create_time" json:"create_time"`
	ModifyTime   *string `gorm:"column:modify_time" json:"modify_time"`
	CreateUser   *string `gorm:"column:create_user" json:"create_user"`
}

// TableName BeanReport's table name
func (*BeanReport) TableName() string {
	return TableNameBeanReport
}

func (*BeanReport) PrimaryKey() []string {
	return []string{"bean_report_id"}
}