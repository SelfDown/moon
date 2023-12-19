package plugins

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	templateService "collect/src/collect/service_imp"
	utils "collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"golang.org/x/crypto/ssh"
	"strings"
)

type Shell struct {
	templateService.BaseHandler
}

func handlerCmd(cmd string, client *ssh.Client, params map[string]interface{}) (result string, err error) {
	// 这里可以进行配置目标执行命令的校验
	session, error := client.NewSession()
	if error != nil {
		return "", error
	}
	defer session.Close()

	bytes, error := session.CombinedOutput(cmd)
	if error != nil {
		return "", error
	}
	result = gocast.ToString(bytes)
	return result, nil
}
func (si *Shell) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	client := ts.GetThirdData(SshName).(*ssh.Client)
	field := handlerParam.Field
	params := template.GetParams()
	var result interface{}

	if !utils.IsValueEmpty(field) {
		// 转成模板
		tpl, error := config.CastTemplate(field)
		if error != nil {
			return common.NotOk(error.Error())
		}
		cmd := utils.RenderTplData(tpl, params).(string)
		// 执行命令
		res, error := handlerCmd(cmd, client, params)
		if error != nil {
			return common.NotOk(error.Error())
		}
		result = res

	}
	fields := handlerParam.Fields
	if len(fields) > 0 {
		cmds := make([]string, len(fields))
		// 依次编译模板
		for index, subField := range fields {
			tpl, error := config.CastTemplate(subField.Field)
			if error != nil {
				return common.NotOk(error.Error())
			}
			cmd := utils.RenderTplData(tpl, params).(string)
			cmds[index] = cmd
		}
		cmd := strings.Join(cmds, "\n")
		res, error := handlerCmd(cmd, client, params)
		if error != nil {
			return common.NotOk(error.Error())
		}
		if utils.IsValueEmpty(fields[0].SaveField) {
			result = res
		} else { // 如果每个子项都设置save_field
			r := make(map[string]interface{})
			tmpArr := strings.Split(res, "\n")
			for index, subField := range fields {
				saveField := subField.SaveField
				if len(tmpArr) > index {
					r[saveField] = tmpArr[index]
				}

			}
			result = r

		}

	}
	r := common.Ok(result, "处理参数成功")
	return r
}
