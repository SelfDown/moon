// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLdapGroup = "ldap_group"

// LdapGroup mapped from table <ldap_group>
type LdapGroup struct {
	LdapGroupID string  `gorm:"column:ldap_group_id;primaryKey" json:"ldap_group_id"`
	Name        *string `gorm:"column:name" json:"name"`
}

// TableName LdapGroup's table name
func (*LdapGroup) TableName() string {
	return TableNameLdapGroup
}

func (*LdapGroup) PrimaryKey() []string {
	return []string{"ldap_group_id"}
}