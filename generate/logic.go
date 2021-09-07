package generate

import (
	"fmt"
	"github.com/SocksStack/vk/utils"
	"path"
	"path/filepath"
	"strings"
)

func CreateLogic(curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	template := strings.Replace(logicTemplate, "{:LOGIC:}", utils.UcWorld(filepath.Base(filename)), -1)
	template = strings.Replace(template, "{:PACKAGE:}", filepath.Base(filepath.Dir(main)), -1)
	utils.CreateFile(main, template)
}

var logicTemplate = `package {:PACKAGE:}

type {:LOGIC:} struct {}

func (i *{:LOGIC:}) {:LOGIC:}() string {
	return "" 
}
`