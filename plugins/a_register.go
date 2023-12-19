package plugins

import (
	templateService "collect/src/collect/service_imp"
)

func GetRegisterList() []templateService.ModuleResult {
	l := make([]templateService.ModuleResult, 0)
	l = append(l, &AnalysisIp{})
	l = append(l, &Shell{})
	// 执行shell
	l = append(l, &Ssh{})
	return l
}
