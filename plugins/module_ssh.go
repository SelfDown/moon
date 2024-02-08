package plugins

import (
	"fmt"
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"golang.org/x/crypto/ssh"
)

type Ssh struct {
	templateService.BaseHandler
	templateService.BaseFlow
}

// close 关闭
func close(_ *config.Template, ts *templateService.TemplateService) {
	// 关闭ssh
	if ts.HasThirdData(SshName) {
		client := ts.GetThirdData(SshName).(*ssh.Client)
		client.Close()
		ts.RemoveThirdData(SshName)
	}

}
func (s *Ssh) Result(template *config.Template, ts *templateService.TemplateService) *common.Result {
	connect := template.DataJsonConfig.Connect
	params := template.GetParams()

	user := utils.RenderVar(gocast.ToString(connect["user"]), params)
	if utils.IsValueEmpty(user) {
		return common.NotOk("用户不能为空")
	}

	password := utils.RenderVar(gocast.ToString(connect["password"]), params)
	if utils.IsValueEmpty(password) {
		return common.NotOk("密码不能为空")
	}
	serverIp := utils.RenderVar(gocast.ToString(connect["server_ip"]), params)
	if utils.IsValueEmpty(serverIp) {
		return common.NotOk("IP不能为空")
	}
	port := utils.RenderVar(gocast.ToString(connect["port"]), params)
	if utils.IsValueEmpty(port) {
		return common.NotOk("端口不能为空")
	}

	config := &ssh.ClientConfig{
		User: user.(string),
		Auth: []ssh.AuthMethod{
			ssh.Password(password.(string)),
		},
		Config: ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%s", serverIp, gocast.ToString(port))

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return common.NotOk(err.Error())
	}
	ts.SetThirdData(SshName, client)

	defer close(template, ts)
	// 流程返回结果
	flowResult := s.Flow(template, ts, templateService.BaseHandlerNode)
	return flowResult
}
