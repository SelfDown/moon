package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type loginInfo struct {
	Address  string `json:"ipaddress"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

//将字符串转换为int
func toInt(str string) int {
	data, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

//初始化ws，将http协议提升为ws协议
func InitialWS(c *gin.Context) (*websocket.Conn, error) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//将http协议提升为ws
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ws, err
}

//实现write方法，保存ws连接
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

type sshConnect struct {
	connect   *ssh.Client    //ssh连接
	session   *ssh.Session   //ssh会话
	stdinPipe io.WriteCloser //标准输入管道
}

// 初始化ssh连接
func (t *sshConnect) InitialSSH(login loginInfo) error {
	//初始化ssh登陆配置
	config := &ssh.ClientConfig{
		User:    login.UserName,
		Auth:    []ssh.AuthMethod{ssh.Password(login.PassWord)},
		Timeout: 5 * time.Second,
		Config: ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	//创建ssh连接
	sshConnect, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", login.Address, login.Port), config)
	if err != nil {
		return err
	}
	t.connect = sshConnect
	return nil
}

//初始化终端
func (t *sshConnect) InitialTerminal(ws *websocket.Conn, rows, cols int) error {
	session, err := t.connect.NewSession()
	if err != nil {
		log.Println(err)
		return err
	}
	t.session = session
	t.stdinPipe, _ = session.StdinPipe()

	wsOutput := WSoutput{
		ws: ws,
	}
	//ssh.stdout and stderr will write output into comboWriter
	session.Stdout = &wsOutput
	session.Stderr = &wsOutput

	//定义terminal模式
	modes := ssh.TerminalModes{
		ssh.ECHO:          1, //开启回显
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	//将pty与远程的会话关联起来
	if err := session.RequestPty("xterm", rows, cols, modes); err != nil {
		return err
	}
	// Shell starts a login shell on the remote host
	if err := session.Shell(); err != nil {
		return err
	}
	return nil
}

//处理连接
func (t *sshConnect) Connect(ws *websocket.Conn) {
	stopCh := make(chan struct{}, 1)

	//启用协程获取webTerminal的输入
	go func() {
		for {
			// 读取web socket信息，msg为用户输入的信息
			_, msg, err := ws.ReadMessage()

			if err != nil {
				log.Println(err)
				stopCh <- struct{}{}
			}
			//心跳检测
			if string(msg) == "ping" {
				continue
			}
			//terminal窗口调整
			if strings.Contains(string(msg), "resize") {
				resizeSlice := strings.Split(string(msg), ":")
				rows, _ := strconv.Atoi(resizeSlice[1])
				cols, _ := strconv.Atoi(resizeSlice[2])
				err := t.session.WindowChange(rows, cols)
				if err != nil {
					log.Println(err)
					stopCh <- struct{}{}
				}
				continue
			}
			_, err = t.stdinPipe.Write(msg)
			if err != nil {
				stopCh <- struct{}{}
				return
			}
		}
	}()

	timer := time.NewTimer(time.Minute * 30)

	defer func() {
		ws.Close()
		t.stdinPipe.Close()
		t.session.Close()
		t.connect.Close()
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	for {
		select {
		case <-stopCh:
			return
		case <-timer.C:
			return
		}
	}
}

func sshTerminal(c *gin.Context) {
	//host := c.DefaultQuery("host", "192.168.232.129")
	cols := c.DefaultQuery("cols", "150")
	rows := c.DefaultQuery("rows", "35")

	sshInfo := loginInfo{
		Address:  "192.168.232.129",
		Port:     22,
		UserName: "root",
		PassWord: "zhang@888",
	}

	//将http协议提升为ws协议
	ws, err := InitialWS(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	connectAction := sshConnect{}
	err = connectAction.InitialSSH(sshInfo)
	err = connectAction.InitialTerminal(ws, toInt(rows), toInt(cols))
	if err != nil {
		ws.WriteMessage(1, []byte(err.Error()))
		ws.Close()
		log.Println(err)
		return
	}
	connectAction.Connect(ws)
}

func main12() {
	r := gin.Default()
	r.GET("/terminal", sshTerminal)
	r.Run(":9090")
}
