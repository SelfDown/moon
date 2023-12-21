package main

import (
	gen "collect/gen"
	templateService "collect/src/collect/service_imp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"moon/model"
	"moon/plugins"
)

func main() {
	gen.GenModel()
}
func main1() {

	// todo go profile 使用
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 生成cookies
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session_id", store))
	r.Static("/static", "./static")
	// 设置数据库
	templateService.SetDatabaseModel(&model.TableData{})
	// 设置外部处理器
	templateService.SetOuterModuleRegister(plugins.GetRegisterList())
	// 添加定时任务
	templateService.RunScheduleService()
	// 添加启动服务
	templateService.RunStartupService()
	r.POST("/template_data/data", func(c *gin.Context) {
		templateService.HandlerRequest(c)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
