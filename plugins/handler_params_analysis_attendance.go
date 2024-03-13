package plugins

import (
	"fmt"
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"time"
)

type AnalysisAttendance struct {
	templateService.BaseHandler
}

func handlerAttendance(checkin map[string]interface{}, dict map[string]map[string]interface{}) {

	schCheckinTime := checkin["sch_checkin_time"].(int64)
	checkin_type := checkin["checkin_type"]
	fmt.Println(checkin_type)
	t := time.Unix(schCheckinTime, 0)
	schDay := utils.DateFormatDay(t)
	var item map[string]interface{}

	if target, ok := dict[schDay]; ok {
		item = target
	} else {
		item = make(map[string]interface{})
	}
	dict[schDay] = item
}

func (si *AnalysisAttendance) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	params := template.GetParams()
	dataList, errMsg := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	attendanceMap := make(map[string]map[string]interface{}, 0)
	for _, checkin := range dataList {
		// 应打卡时间
		fmt.Println(checkin)

	}
	target := make([]map[string]interface{}, 0)
	for _, item := range attendanceMap {
		target = append(target, item)
	}

	r := common.Ok(target, "处理参数成功")
	return r
}
