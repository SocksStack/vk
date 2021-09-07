package generate

import (
	"github.com/SocksStack/vk/utils"
	"path"
	"strings"
)

func CreateReadMe(moduleName, curPath string)  {
	main := path.Join(curPath, "README.md")
	template := strings.Replace(readmeTemplate, "{:MODULE:}", moduleName, -1)
	utils.CreateFile(main, template)
}

var readmeTemplate = "# {:MODULE:}\n\n#### 介绍\n\nxxx\n\n#### 软件架构\n软件架构说明\n\n#### 项目结构\n\n```\ndemo\n ├── bootstrap\n │   └── bootstrap.go\n ├── config\n │   ├── default.yml\n │   ├── development.yml\n │   ├── production.yml\n │   └── test.yml\n ├── handler\n │   └── index.go\n ├── middleware\n │   └── init.go\n ├── router\n │   └── router.go\n ├── main.go\n ├── go.mod\n ├── go.sum\n └── README.md\n```\n\n#### 安装教程\n\n1.  xxxx\n2.  xxxx\n3.  xxxx\n\n#### 使用说明\n\n1.  xxxx\n2.  xxxx\n3.  xxxx\n\n#### 参与贡献\n\n1.  Fork 本仓库\n2.  新建 Feat_xxx 分支\n3.  提交代码\n4.  新建 Pull Request\n\n\n#### 其他说明\n\n1.本项目由 [SocksStack/vk](https://github.com/SocksStack/vk) 生成\n"