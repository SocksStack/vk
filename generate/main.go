package generate

import (
	"fmt"
	"github.com/SocksStack/vk/utils"
	"path"
	"strings"
)

func CreateMain(moduleName, curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	template := strings.Replace(mainTemplate, "{:MODULE:}", moduleName, -1)
	utils.CreateFile(main, template)
}

var mainTemplate = `package main

import (
	"{:MODULE:}/bootstrap"
)

func main() {
	bootstrap.Run()
}
`
