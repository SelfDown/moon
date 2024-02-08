package plugins

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"reflect"
)

// MultiArr 数组乘以数组
type MultiArr struct {
	templateService.BaseHandler
}

func ArrToArrMap(data interface{}, field string) ([]map[string]interface{}, string) {

	original := reflect.ValueOf(data)
	dataList := make([]map[string]interface{}, original.Len())
	for i := 0; i < original.Len(); i++ {

		dataItem := make(map[string]interface{})
		value := original.Index(i).Interface()
		dataItem[field] = value
		dataList[i] = dataItem
	}
	return dataList, ""
}

func (si *MultiArr) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	params := template.GetParams()
	arr, errMsg := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(handlerParam.Foreach + errMsg)
	}
	rightObj := utils.RenderVar(handlerParam.Right, params)
	if rightObj == nil {
		return common.NotOk(handlerParam.Right + "数据为空")
	}
	rightArr, errMsg := ArrToArrMap(rightObj, handlerParam.RightField)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(handlerParam.Right + errMsg)
	}
	fSize := len(arr)
	// 一次性初始化大小
	target := make([]map[string]interface{}, fSize*len(rightArr))
	// 右边的参数字段，重复利用了
	for i, fItem := range arr { // 编辑foreach 数组
		for j, rItem := range rightArr { // 将右边的数据复制到左边上面去
			targetItem := utils.Copy(fItem).(map[string]interface{})

			for _, field := range handlerParam.Fields { // 挨个字段复制
				//  根据template 渲染值
				targetItem[field.Field] = utils.RenderTplData(field.TemplateTpl, rItem)
			}
			// 定位到数组
			target[fSize*i+j] = targetItem
		}

	}

	r := common.Ok(target, "处理参数成功")
	return r
}
