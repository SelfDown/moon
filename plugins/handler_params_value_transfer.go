package plugins

import (
	"encoding/json"
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"strings"
)

type ValueTransfer struct {
	templateService.BaseHandler
}
type TransferValueData struct {
	Fields []Rule
}
type Rule struct {
	Field string
	Rules map[string]interface{}
}

func (si *ValueTransfer) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	params := template.GetParams()
	field := handlerParam.Field
	fieldValue := utils.RenderVar(field, params)
	if utils.IsValueEmpty(fieldValue) {
		return common.NotOk("来源字段不存在")
	}
	if utils.IsValueEmpty(template.FileData) {
		return common.NotOk("data_file 字段不存在")
	}
	t := TransferValueData{}
	json.Unmarshal([]byte(template.FileData), &t)
	target := fieldValue.(map[string]interface{})
	for _, fieldRule := range t.Fields {
		fieldName := fieldRule.Field
		val, ok := target[fieldName]
		if !ok {
			continue
		}
		if utils.IsValueEmpty(val) {
			continue
		}
		varStr := val.(string)
		rules := fieldRule.Rules
		for rule := range rules {
			if strings.HasPrefix(rule, "~") && strings.HasSuffix(rule, "~") { // 中间包含
				tag := rule[1 : len(rule)-1]
				if strings.Contains(varStr, tag) {
					target[fieldName] = rules[rule]
					continue
				}
			} else if strings.HasPrefix(rule, "~") { // 以xx 结尾
				tag := rule[1:]
				if strings.HasSuffix(varStr, tag) {
					target[fieldName] = rules[rule]
					continue
				}
			} else if strings.HasSuffix(rule, "~") { // 以xx开头
				tag := rule[0 : len(rule)-1]
				if strings.HasPrefix(varStr, tag) {
					target[fieldName] = rules[rule]
					continue
				}
			} else { // 精确匹配
				tag := rule
				if tag == varStr {
					target[fieldName] = rules[rule]
					continue
				}
			}
		}

	}

	r := common.Ok(target, "处理参数成功")
	return r
}
