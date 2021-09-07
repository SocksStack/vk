package generate

import (
	"fmt"
	"github.com/SocksStack/vk/utils"
"path"
"strings"
)

func CreateRouter(moduleName, curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	template := strings.Replace(routerTemplate, "{:MODULE:}", moduleName, -1)
	utils.CreateFile(main, template)
}

var routerTemplate = `package router

import (
	"{:MODULE:}/handler"
	"github.com/SocksStack/common/constract"
)

func Routes() constract.Routes {
	return constract.Routes{
		&handler.Index{},
	}
}

`
