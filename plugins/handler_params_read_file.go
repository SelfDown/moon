package plugins

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type ReadFile struct {
	templateService.BaseHandler
}

func (si *ReadFile) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	params := template.GetParams()
	valTpl := handlerParam.ValueTpl
	path := utils.RenderTplData(valTpl, params).(string)
	content, success := utils.ReadFileContent(path)
	if !success {
		common.NotOk(path + "文件不存在")
	}
	r := common.Ok(content, "处理参数成功")
	return r
}
