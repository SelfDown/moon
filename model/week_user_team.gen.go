// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWeekUserTeam = "week_user_team"

// WeekUserTeam mapped from table <week_user_team>
type WeekUserTeam struct {
	WeekUserTeamID string  `gorm:"column:week_user_team_id;primaryKey" json:"week_user_team_id"`
	BeanReportID   *string `gorm:"column:bean_report_id" json:"bean_report_id"`
	Username       string  `gorm:"column:username;not null" json:"username"`
	Nick           string  `gorm:"column:nick;not null" json:"nick"`
	Team           string  `gorm:"column:team;not null" json:"team"`
}

// TableName WeekUserTeam's table name
func (*WeekUserTeam) TableName() string {
	return TableNameWeekUserTeam
}

func (*WeekUserTeam) PrimaryKey() []string {
	return []string{"week_user_team_id"}
}