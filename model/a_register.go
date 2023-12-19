package model

import (
	templateService "collect/src/collect/service_imp"
	utils "collect/src/collect/utils"
)

var modelMap map[string]interface{}
var primaryKeyMap map[string][]string

type TableData struct {
	templateService.DatabaseModel
}

// 生成一个脚本自动填充这个
func init() {
	//todo 如果用hashmap 效率慢，可以换二叉树，目前1、200个表很快
	modelMap = make(map[string]interface{})
	primaryKeyMap = make(map[string][]string)
	addTable()
}
func (*TableData) GetModel(tableName string) interface{} {
	return modelMap[tableName]
}
func (*TableData) CloneModel(tableName string) interface{} {
	data := modelMap[tableName]
	return utils.Copy(data)
}
func (*TableData) GetPrimaryKey(tableName string) []string {
	return primaryKeyMap[tableName]
}
