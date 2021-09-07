package generate

import (
	"fmt"
	"github.com/SocksStack/vk/utils"
	"path"
	"path/filepath"
	"strings"
)

func CreateService(curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	template := strings.Replace(serviceTemplate, "{:SERVICE:}", utils.UcWorld(path.Base(filename)), -1)
	template = strings.Replace(template, "{:PACKAGE:}", filepath.Base(filepath.Dir(main)), -1)
	utils.CreateFile(main, template)
}

var serviceTemplate = `package {:PACKAGE:}

type {:SERVICE:} struct {}

func (i *{:SERVICE:}) {:SERVICE:}() string {
	return "" 
}
`