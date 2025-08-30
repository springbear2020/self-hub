package email

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/plugin/email/global"
	"github.com/springbear2020/self-hub/server/plugin/email/router"
)

type emailPlugin struct{}

func CreateEmailPlug(To, From, Host, Secret, Nickname string, Port int, IsSSL bool, IsLoginAuth bool) *emailPlugin {
	global.GlobalConfig.To = To
	global.GlobalConfig.From = From
	global.GlobalConfig.Host = Host
	global.GlobalConfig.Secret = Secret
	global.GlobalConfig.Nickname = Nickname
	global.GlobalConfig.Port = Port
	global.GlobalConfig.IsSSL = IsSSL
	global.GlobalConfig.IsLoginAuth = IsLoginAuth
	return &emailPlugin{}
}

func (*emailPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitEmailRouter(group)
}

func (*emailPlugin) RouterPath() string {
	return "email"
}
