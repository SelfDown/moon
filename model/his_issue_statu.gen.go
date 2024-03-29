// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHisIssueStatu = "his_issue_statu"

// HisIssueStatu mapped from table <his_issue_statu>
type HisIssueStatu struct {
	HisIssueStatuID   string  `gorm:"column:his_issue_statu_id;primaryKey" json:"his_issue_statu_id"`
	StatuText         string  `gorm:"column:statu_text;not null" json:"statu_text"`
	Notes             *string `gorm:"column:notes" json:"notes"`
	Weight            *int32  `gorm:"column:weight" json:"weight"`
	OrderIndex        *int32  `gorm:"column:order_index" json:"order_index"`
	NotifyURL         *string `gorm:"column:notify_url" json:"notify_url"`
	StatuTextCustomer *string `gorm:"column:statu_text_customer" json:"statu_text_customer"`
}

// TableName HisIssueStatu's table name
func (*HisIssueStatu) TableName() string {
	return TableNameHisIssueStatu
}

func (*HisIssueStatu) PrimaryKey() []string {
	return []string{"his_issue_statu_id"}
}