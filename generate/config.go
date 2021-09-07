package generate

import (
	"github.com/SocksStack/vk/utils"
	"path"
	"strings"
)

func CreateConfig(curPath string)  {
	defaultConfig := path.Join(curPath, "default.yml")
	developmentConfig := path.Join(curPath, "development.yml")
	productionConfig := path.Join(curPath, "production.yml")
	testConfig := path.Join(curPath, "test.yml")
	utils.CreateFile(defaultConfig, strings.Replace(strings.Replace(defaultConfigTemplate, "{:ACTIVE:}", "development", -1), "{:MODE:}", "debug", -1))
	utils.CreateFile(developmentConfig, strings.Replace(strings.Replace(defaultConfigTemplate, "{:ACTIVE:}", "development", -1), "{:MODE:}", "debug", -1))
	utils.CreateFile(productionConfig, strings.Replace(strings.Replace(defaultConfigTemplate, "{:ACTIVE:}", "production", -1), "{:MODE:}", "release", -1))
	utils.CreateFile(testConfig, strings.Replace(strings.Replace(defaultConfigTemplate, "{:ACTIVE:}", "test", -1), "{:MODE:}", "test", -1))
}

var defaultConfigTemplate = `server:
  host: 0.0.0.0
  port: 8080
  mode: {:MODE:} #运行环境 debug, release, test
active: {:ACTIVE:} #配置环境 development, production, test
`
