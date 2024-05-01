package plugins

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"time"
)

type AnalysisAttendance struct {
	templateService.BaseHandler
}

/**
@param checkin 一次考勤记录
@dict 按照天的考勤
*/
const CheckInTypeOnDuty = "上班打卡"
const CheckInTypeOffDuty = "下班打卡"

// 缺卡
const NoCheckIn = "no_checkin"

// 正常打卡
const NormalCheckIn = "normal"

// 非正常
const NotNormalCheckIn = "not_normal"

// 迟到
const Late = "late"

// 早退
const Early = "early"

// 旷工
const NoAttendance = "no_attendance"

func getRuleTime(ruleList []map[string]interface{}, ruleType string) string {
	var weight int64
	var ruleTime string
	for _, item := range ruleList {
		configRuleType := gocast.ToString(item["rule_type"])
		if configRuleType == ruleType {
			w := gocast.ToInt64(item["weight"])
			if w > weight {
				weight = w
				ruleTime = gocast.ToString(item["rule_time"])
			}
		}
	}
	return ruleTime
}
func getAttendanceType(checkin map[string]interface{}, checkinType string, params map[string]interface{}) string {
	exceptionType := checkin["exception_type"]
	attendanceTypeData := ""
	// todo 这里要处理成上班迟到、下班早退。这里先预留
	ruleList, _ := utils.RenderVarToArrMap("rule_list", params)
	if utils.IsValueEmpty(exceptionType) {
		attendanceTypeData = NormalCheckIn
	} else if exceptionType == "时间异常" {

		checkinTime := checkin["checkin_time"].(int64)
		schCheckinTime := checkin["sch_checkin_time"].(int64)
		checkinTimeTmp := time.Unix(schCheckinTime, 0)
		schDay := utils.DateFormatDay(checkinTimeTmp)
		if checkinType == CheckInTypeOnDuty { // 上班打卡类型

			ruleTime := getRuleTime(ruleList, "morning")

			if utils.IsValueEmpty(ruleTime) {
				panic("未设置上班截止时间")
			}
			stopTime := utils.DateTimeUnix(schDay + " " + ruleTime)
			if checkinTime > stopTime { // 如果打卡时间大于截止时间
				return Late
			} else {
				return NormalCheckIn
			}

		} else if checkinType == CheckInTypeOffDuty {
			ruleTime := getRuleTime(ruleList, "afternoon")

			if utils.IsValueEmpty(ruleTime) {
				panic("未设置下班最早打卡时间")
			}
			startTime := utils.DateTimeUnix(schDay + " " + ruleTime)
			if startTime > checkinTime { // 早退
				return Early
			} else {
				return NormalCheckIn
			}
		}
	} else if exceptionType == "未打卡" {
		attendanceTypeData = "no_checkin"
	} else {
		attendanceTypeData = "unknown"
	}
	return attendanceTypeData
}

type AttendanceHandler interface {
	handler(checkin map[string]interface{}, dict map[string]map[string]interface{}, params map[string]interface{}) (map[string]interface{}, string)
}

func getAttendDayItem(checkin map[string]interface{}, dict map[string]map[string]interface{}) (map[string]interface{}, string, string) {
	schCheckinTime := checkin["sch_checkin_time"].(int64)
	if utils.IsValueEmpty(checkin["sch_checkin_time"]) { // 如果没有应考勤日期，直接取考勤日期
		schCheckinTime = checkin["checkin_time"].(int64)
	}
	checkinTime := time.Unix(schCheckinTime, 0)
	schDay := utils.DateFormatDay(checkinTime)
	var item map[string]interface{}
	if target, ok := dict[schDay]; ok {
		item = target
	} else {
		item = make(map[string]interface{})
	}
	checkinType := gocast.ToString(checkin["checkin_type"])
	return item, schDay, checkinType
}

type DayHandler struct {
	AttendanceHandler
}

// 处理应该打卡天
func (t *DayHandler) handler(checkin map[string]interface{}, dict map[string]map[string]interface{}, params map[string]interface{}) (map[string]interface{}, string) {

	item, schDay, _ := getAttendDayItem(checkin, dict)
	item["attendance_day"] = schDay

	return item, schDay
}

type UserHandler struct {
	AttendanceHandler
}

// 处理用户
func (t *UserHandler) handler(checkin map[string]interface{}, dict map[string]map[string]interface{}, params map[string]interface{}) (map[string]interface{}, string) {

	item, schDay, _ := getAttendDayItem(checkin, dict)
	userDict := params["user_dict"].(map[string]interface{})
	userid := gocast.ToString(checkin["userid"])
	item["user_id"] = userDict[userid]
	return item, schDay
}

// 处理上下班考勤类型
type AttendanceTimeHandler struct {
	AttendanceHandler
}

func (t *AttendanceTimeHandler) handler(checkin map[string]interface{}, dict map[string]map[string]interface{}, params map[string]interface{}) (map[string]interface{}, string) {
	// 上班、下班打卡
	item, day, checkinType := getAttendDayItem(checkin, dict)
	checkinTime := checkin["checkin_time"].(int64)
	checkinTimeActual := time.Unix(checkinTime, 0)
	checkFmt := ""
	attendanceType := getAttendanceType(checkin, checkinType, params)
	// 处理没有打开，但是显示打开时间的那种
	if attendanceType != NoCheckIn {
		checkFmt = utils.DateFormatDefault(checkinTimeActual)
	}
	if checkinType == CheckInTypeOnDuty { // 上班打卡类型
		item["attendance_morning_time"] = checkFmt
	} else if checkinType == CheckInTypeOffDuty {
		item["attendance_afternoon_time"] = checkFmt
	}
	return item, day
}

// 处理上下班考勤类型
type AttendanceTypeHandler struct {
	AttendanceHandler
}

func (t *AttendanceTypeHandler) handler(checkin map[string]interface{}, dict map[string]map[string]interface{}, params map[string]interface{}) (map[string]interface{}, string) {
	// 上班、下班打卡
	item, day, checkinType := getAttendDayItem(checkin, dict)
	attendanceTypeData := getAttendanceType(checkin, checkinType, params)
	if checkinType == CheckInTypeOnDuty { // 上班打卡类型
		item["attendance_morning_type"] = attendanceTypeData
	} else if checkinType == CheckInTypeOffDuty {
		item["attendance_afternoon_type"] = attendanceTypeData
	}
	return item, day
}

// 处理上下班考勤类型
type AttendanceLocationHandler struct {
	AttendanceHandler
}

func (t *AttendanceLocationHandler) handler(checkin map[string]interface{}, dict map[string]map[string]interface{}, params map[string]interface{}) (map[string]interface{}, string) {
	// 上班、下班打卡
	item, day, checkinType := getAttendDayItem(checkin, dict)
	locationDetail := checkin["location_detail"]
	if checkinType == CheckInTypeOnDuty { // 上班打卡类型
		item["attendance_morning_location"] = locationDetail
	} else if checkinType == CheckInTypeOffDuty {
		item["attendance_afternoon_location"] = locationDetail
	}
	return item, day
}

// 处理上下班考勤类型
type AttendanceAllDayHandler struct {
	AttendanceHandler
}

func (t *AttendanceAllDayHandler) handler(checkin map[string]interface{}, dict map[string]map[string]interface{}, params map[string]interface{}) (map[string]interface{}, string) {
	// 上班、下班打卡
	item, day, _ := getAttendDayItem(checkin, dict)
	attendanceMorningType := item["attendance_morning_type"]
	attendanceAfternoonType := item["attendance_afternoon_type"]
	if attendanceMorningType == nil { // 如果没有数据就是没有打开
		attendanceMorningType = NoCheckIn
	}
	if attendanceAfternoonType == nil {
		attendanceAfternoonType = NoCheckIn
	}
	if attendanceMorningType == NormalCheckIn && attendanceAfternoonType == NormalCheckIn {
		// 正常打卡
		item["attendance_type"] = NormalCheckIn
	} else if attendanceMorningType == NoCheckIn && attendanceAfternoonType == NoCheckIn {
		// 上午和下午都没有打开
		item["attendance_type"] = NoAttendance
	} else if attendanceMorningType != NormalCheckIn && attendanceAfternoonType != NormalCheckIn {
		// 上午和下午都不正常
		item["attendance_type"] = NotNormalCheckIn
	} else if attendanceMorningType == NormalCheckIn {
		// 上午正常打卡，下午没有
		item["attendance_type"] = attendanceAfternoonType
	} else if attendanceAfternoonType == NormalCheckIn {
		// 下午正常打卡，上午没有
		item["attendance_type"] = attendanceMorningType
	} else {
		// 如果是当天打开，下午还没有打卡
		item["attendance_type"] = attendanceMorningType
	}
	return item, day
}

func (si *AnalysisAttendance) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	params := template.GetParams()
	dataList, errMsg := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	//attendanceMap := make(map[string]map[string]interface{}, 0)
	//用户#天#考勤数据
	attendanceMapAll := make(map[string]map[string]map[string]interface{}, 0)
	attendanceHandlerList := make([]AttendanceHandler, 0)
	// 处理用户ID
	attendanceHandlerList = append(attendanceHandlerList, &UserHandler{})
	// 处理考勤日
	attendanceHandlerList = append(attendanceHandlerList, &DayHandler{})

	// 处理时间
	attendanceHandlerList = append(attendanceHandlerList, &AttendanceTimeHandler{})
	// 处理地点
	attendanceHandlerList = append(attendanceHandlerList, &AttendanceLocationHandler{})
	// 处理类型
	attendanceHandlerList = append(attendanceHandlerList, &AttendanceTypeHandler{})
	// 处理天类型
	attendanceHandlerList = append(attendanceHandlerList, &AttendanceAllDayHandler{})
	//todo 处理补卡的
	userList := params["user_list"].([]map[string]interface{})
	userDict := make(map[string]interface{})
	for _, user := range userList {
		wechatUserid := gocast.ToString(user["work_number"])
		userDict[wechatUserid] = user["userid"]
	}
	params["user_dict"] = userDict
	for _, checkin := range dataList {
		// 处理考勤
		userid := gocast.ToString(checkin["userid"])
		attendanceMap, ok := attendanceMapAll[userid]
		if !ok {
			attendanceMap = make(map[string]map[string]interface{}, 0)
		}
		for _, handler := range attendanceHandlerList {
			item, day := handler.handler(checkin, attendanceMap, params)
			attendanceMap[day] = item
		}
		attendanceMapAll[userid] = attendanceMap

	}
	target := make([]map[string]interface{}, 0)
	for _, item := range attendanceMapAll {
		for _, second := range item {
			target = append(target, second)
		}

	}

	r := common.Ok(target, "处理参数成功")
	return r
}
