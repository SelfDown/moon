// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSoftModelNew = "soft_model_new"

// SoftModelNew mapped from table <soft_model_new>
type SoftModelNew struct {
	SoftModelName *string `gorm:"column:soft_model_name" json:"soft_model_name"`
	SoftType      *string `gorm:"column:soft_type" json:"soft_type"`
	SoftDataJSON  *string `gorm:"column:soft_data_json" json:"soft_data_json"`
	UUIDID        string  `gorm:"column:uuid_id;primaryKey" json:"uuid_id"`
	IsDelete      *string `gorm:"column:is_delete" json:"is_delete"`
	OpUser        *string `gorm:"column:op_user" json:"op_user"`
}

// TableName SoftModelNew's table name
func (*SoftModelNew) TableName() string {
	return TableNameSoftModelNew
}

func (*SoftModelNew) PrimaryKey() []string {
	return []string{"uuid_id"}
}