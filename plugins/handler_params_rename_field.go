package plugins

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	"regexp"
)

type RenameField struct {
	templateService.BaseHandler
}

func (si *RenameField) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	params := template.GetParams()
	field := handlerParam.Field
	for key, value := range params {
		re := regexp.MustCompile(field)
		found := re.FindAllStringSubmatch(key, -1)
		if found == nil {
			continue
		}
		fieldName := found[0][1]
		delete(params, key)
		params[fieldName] = value

	}

	r := common.Ok(nil, "处理参数成功")
	return r
}
