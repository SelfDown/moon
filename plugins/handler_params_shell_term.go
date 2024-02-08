package plugins

import (
	"encoding/json"
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
	"time"
	"unicode/utf8"
)

type ShellTerm struct {
	templateService.BaseHandler
}

var upgrader = websocket.Upgrader{

	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options
type WSoutput struct {
	ws *websocket.Conn
}

// Write: implement Write interface to write bytes from ssh server into bytes.Buffer.
func (w *WSoutput) Write(p []byte) (int, error) {
	// 处理非utf8字符
	if !utf8.Valid(p) {
		bufStr := string(p)
		buf := make([]rune, 0, len(bufStr))
		for _, r := range bufStr {
			if r == utf8.RuneError {
				buf = append(buf, []rune("@")...)
			} else {
				buf = append(buf, r)
			}
		}
		p = []byte(string(buf))
	}
	err := w.ws.WriteMessage(websocket.TextMessage, p)
	return len(p), err
}
func updateTimer(timer *time.Timer) {
	timer.Reset(time.Minute * 20)

}

type MessageData struct {
	Type string
	Rows int
	Cols int
	Data string
}

func (si *ShellTerm) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	client := ts.GetThirdData(SshName).(*ssh.Client)
	var result interface{}
	context := ts.GetContext()
	w := context.Writer
	req := context.Request
	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		//log.Print("upgrade:", err)
		return common.NotOk(err.Error())
	}

	session, error := client.NewSession()
	if error != nil {
		return common.NotOk(error.Error())
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := session.RequestPty("linux", 80, 40, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	wsOutput := WSoutput{
		ws: c,
	}
	session.Stdout = &wsOutput
	session.Stderr = &wsOutput
	pipeIn, _ := session.StdinPipe()
	defer func() {
		// 输入通道关闭
		pipeIn.Close()
		//连接回话关闭
		session.Close()
		// ssh 客户端关闭
		client.Close()
		// websocket 关闭
		c.Close()

	}()
	session.Shell()
	// 如果30分钟内容不输入，直接关闭
	timer := time.NewTimer(time.Minute * 20)
	updateTime := utils.Debounce(func() {
		updateTimer(timer)
	}, 5)
	stopCh := make(chan struct{}, 1)
	go func() {
		for {
			_, data, err := c.ReadMessage()
			messageData := MessageData{}
			json.Unmarshal(data, &messageData)
			if err != nil {
				log.Println("read:", err)
				stopCh <- struct{}{}
				break
			}
			updateTime()
			if messageData.Type == "text" {
				_, err = pipeIn.Write([]byte(messageData.Data))
			} else if messageData.Type == "resize" {
				session.WindowChange(messageData.Rows, messageData.Cols)
			}
			if err != nil {
				log.Println(err)
				stopCh <- struct{}{}
				break
			}
		}
	}()
	// 如果停止了、或者超时了，直接返回
	for {
		select {
		case <-stopCh:
			return common.Ok(result, "结束执行")
		case <-timer.C:
			return common.Ok(result, "执行超时")
		}
	}

	return common.Ok(result, "处理参数成功")
}
