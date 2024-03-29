// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDjangoApschedulerDjangojobexecution = "django_apscheduler_djangojobexecution"

// DjangoApschedulerDjangojobexecution mapped from table <django_apscheduler_djangojobexecution>
type DjangoApschedulerDjangojobexecution struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Status    string    `gorm:"column:status;not null" json:"status"`
	RunTime   time.Time `gorm:"column:run_time;not null" json:"run_time"`
	Duration  *float64  `gorm:"column:duration" json:"duration"`
	Started   *float64  `gorm:"column:started" json:"started"`
	Finished  *float64  `gorm:"column:finished" json:"finished"`
	Exception *string   `gorm:"column:exception" json:"exception"`
	Traceback *string   `gorm:"column:traceback" json:"traceback"`
	JobID     int32     `gorm:"column:job_id;not null" json:"job_id"`
}

// TableName DjangoApschedulerDjangojobexecution's table name
func (*DjangoApschedulerDjangojobexecution) TableName() string {
	return TableNameDjangoApschedulerDjangojobexecution
}

func (*DjangoApschedulerDjangojobexecution) PrimaryKey() []string {
	return []string{"id"}
}