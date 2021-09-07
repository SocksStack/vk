package generate

import (
	"fmt"
	"github.com/SocksStack/vk/utils"
	"path"
	"strings"
)

func CreateBootstrap(moduleName, curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	template := strings.Replace(bootstrapTemplate, "{:MODULE:}", moduleName, -1)
	utils.CreateFile(main, template)
}

var bootstrapTemplate = `package bootstrap

import (
	"{:MODULE:}/middleware"
	"{:MODULE:}/router"
	"github.com/gobuffalo/packr"
	"github.com/SocksStack/common/gin"
	"github.com/SocksStack/common/config"
)

var (
	mode string
	cfg gin.HttpConfig
)

func init() {
	config.Init(packr.NewBox("../config"))
	mode = config.Config.GetString("server.mode")
	cfg = gin.HttpConfig{
		Host: config.Config.GetString("server.host"),
		Port: config.Config.GetInt("server.port"),
	}
}

func Run() {
	gin.Default(mode).SetConfig(cfg).
		Use(middleware.Middlewares()...).
		Route(router.Routes()...).
		Run()
}
`
