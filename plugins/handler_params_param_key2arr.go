package plugins

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"regexp"
)

type ParamKey2Arr struct {
	templateService.BaseHandler
}

func (si *ParamKey2Arr) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	params := template.GetParams()
	field := handlerParam.Field
	dataList := make([]interface{}, 0)
	for key, value := range params {
		re := regexp.MustCompile(field)
		found := re.FindAllStringSubmatch(key, -1)
		if found == nil {
			continue
		}
		fieldName := found[0][1]
		dataItem := make(map[string]interface{})
		dataItem["param_value"] = value
		ok := utils.RenderTplBool(handlerParam.IfTemplateTpl, dataItem)
		if !ok {
			continue
		}
		dataList = append(dataList, fieldName)
	}

	r := common.Ok(dataList, "处理参数成功")
	return r
}
