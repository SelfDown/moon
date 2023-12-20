package model

func addTable() {

	//1 collect_doc
	tableCollectDoc := CollectDoc{}
	modelMap["collect_doc"] = tableCollectDoc
	primaryKeyMap["collect_doc"] = tableCollectDoc.PrimaryKey()
	//2 collect_doc_demo
	tableCollectDocDemo := CollectDocDemo{}
	modelMap["collect_doc_demo"] = tableCollectDocDemo
	primaryKeyMap["collect_doc_demo"] = tableCollectDocDemo.PrimaryKey()
	//3 collect_doc_important
	tableCollectDocImportant := CollectDocImportant{}
	modelMap["collect_doc_important"] = tableCollectDocImportant
	primaryKeyMap["collect_doc_important"] = tableCollectDocImportant.PrimaryKey()
	//4 collect_doc_params
	tableCollectDocParams := CollectDocParams{}
	modelMap["collect_doc_params"] = tableCollectDocParams
	primaryKeyMap["collect_doc_params"] = tableCollectDocParams.PrimaryKey()
	//5 collect_doc_result
	tableCollectDocResult := CollectDocResult{}
	modelMap["collect_doc_result"] = tableCollectDocResult
	primaryKeyMap["collect_doc_result"] = tableCollectDocResult.PrimaryKey()
	//6 config_detail
	tableConfigDetail := ConfigDetail{}
	modelMap["config_detail"] = tableConfigDetail
	primaryKeyMap["config_detail"] = tableConfigDetail.PrimaryKey()
	//7 config_group
	tableConfigGroup := ConfigGroup{}
	modelMap["config_group"] = tableConfigGroup
	primaryKeyMap["config_group"] = tableConfigGroup.PrimaryKey()
	//8 doc_group
	tableDocGroup := DocGroup{}
	modelMap["doc_group"] = tableDocGroup
	primaryKeyMap["doc_group"] = tableDocGroup.PrimaryKey()
}