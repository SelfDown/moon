package plugins

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	templateService "collect/src/collect/service_imp"
	utils "collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"strings"
)

type AnalysisIp struct {
	templateService.BaseHandler
}

func (si *AnalysisIp) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	params := template.GetParams()
	field := utils.RenderVar(handlerParam.Field, params).(string)
	arr := strings.Split(field, "-")
	ipList := make([]string, 0)
	start := arr[0]
	//添加开始ip
	ipList = append(ipList, start)
	startTmp := strings.Split(start, ".")
	if len(startTmp) < 4 {
		return common.NotOk(field + "格式有误，必须是xxx.xxx.xxx.xxx 这个格式")
	}
	if gocast.ToInt(startTmp[0]) > 255 ||
		gocast.ToInt(startTmp[1]) > 255 ||
		gocast.ToInt(startTmp[2]) > 255 ||
		gocast.ToInt(startTmp[3]) > 255 {
		return common.NotOk("IP 格式有误,必须小于256")
	}
	// 如果只有一个ip
	if len(arr) == 1 {
		return common.Ok(ipList, "ip 处理成功")
	}

	startIp := gocast.ToInt(startTmp[3])
	// - 后面的数据
	endIp := gocast.ToInt(arr[1])
	for i := startIp + 1; i <= endIp; i++ {
		targetIp := make([]string, 4)
		targetIp[0] = startTmp[0]
		targetIp[1] = startTmp[1]
		targetIp[2] = startTmp[3]
		targetIp[3] = gocast.ToString(i)
		ipTmp := strings.Join(targetIp, ".")
		ipList = append(ipList, ipTmp)
	}

	r := common.Ok(ipList, "处理参数成功")
	return r
}
