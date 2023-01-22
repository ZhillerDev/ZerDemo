package preinit

import (
	"fmt"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/gin-gonic/gin"
)

type machineInfo struct {
	innerIpv4 string
	ipv4List  []string
	macList   []string
}

func ServerInit(engine *gin.Engine) {
	setStaticDocument(engine) // 设置静态文件夹
	getMachineInfo()          // 获取服务器信息
}

func getMachineInfo() {
	infos := machineInfo{
		innerIpv4: netutil.GetInternalIp(),
		ipv4List:  netutil.GetIps(),
		macList:   netutil.GetMacAddrs(),
	}
	fmt.Println(fmt.Sprintf("%+v", infos))
}

func setStaticDocument(engine *gin.Engine) {
	engine.Static("/static", "./runtime")
}
