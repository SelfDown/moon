package main

import (
	gen "github.com/SelfDown/collect/gen"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	collect "github.com/SelfDown/collect/src/collect/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"moon/model"
	"moon/plugins"
	"net/http"
)

func main1() {
	gen.GenModel()
}

func main() {

	// todo go profile 使用
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 生成cookies
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session_id", store))
	r.Static("/static", "./static")
	r.Static("/ssh", "./frontend/ssh")
	r.Static("/cms", "./frontend/cms")

	// 设置数据库
	templateService.SetDatabaseModel(&model.TableData{})
	// 设置外部处理器
	templateService.SetOuterModuleRegister(plugins.GetRegisterList())
	// 添加定时任务
	templateService.RunScheduleService()
	// 添加启动服务
	templateService.RunStartupService()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ssh")
	})
	r.POST("/template_data/data", func(c *gin.Context) {
		templateService.HandlerRequest(c)
	})

	r.GET("/template_data/ws/:token", func(context *gin.Context) {

		templateService.HandlerWsRequest(context)

	})
	serverPort := collect.GetAppKey("server_port")
	r.Run(":" + serverPort) // listen and serve on 0.0.0.0:8080
}
