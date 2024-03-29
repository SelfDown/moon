// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameWebsqlFavorite = "websql_favorite"

// WebsqlFavorite mapped from table <websql_favorite>
type WebsqlFavorite struct {
	WebsqlFavoriteID string     `gorm:"column:websql_favorite_id;primaryKey" json:"websql_favorite_id"`
	FavoriteGroupID  *string    `gorm:"column:favorite_group_id" json:"favorite_group_id"`
	Userid           *string    `gorm:"column:userid" json:"userid"`
	FavoriteName     *string    `gorm:"column:favorite_name" json:"favorite_name"`
	FavoriteType     *string    `gorm:"column:favorite_type" json:"favorite_type"`
	UserID           *string    `gorm:"column:user_id" json:"user_id"`
	SqlText          *string    `gorm:"column:sql_text" json:"sql_text"`
	CreateTime       *time.Time `gorm:"column:create_time" json:"create_time"`
	ModifyTime       *time.Time `gorm:"column:modify_time" json:"modify_time"`
	Comments         *string    `gorm:"column:comments" json:"comments"`
	Orderindex       *int32     `gorm:"column:orderindex" json:"orderindex"`
}

// TableName WebsqlFavorite's table name
func (*WebsqlFavorite) TableName() string {
	return TableNameWebsqlFavorite
}

func (*WebsqlFavorite) PrimaryKey() []string {
	return []string{"websql_favorite_id"}
}