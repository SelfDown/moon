package main

import (
	gen "github.com/SelfDown/collect/gen"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"moon/model"
	"moon/plugins"
	"net/http"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options
func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
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
	r.GET("/template_data/ws/:token", func(context *gin.Context) {
		templateService.HandlerWsRequest(context)
		//ws(context.Writer, context.Request)
	})
	r.Run(":8009") // listen and serve on 0.0.0.0:8080
}
