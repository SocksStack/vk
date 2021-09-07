package generate

import (
	"fmt"
	"github.com/SocksStack/vk/utils"
	"path"
	"path/filepath"
	"strings"
)

func CreateRepository(curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	template := strings.Replace(repositoryTemplate, "{:REPOSITORY:}", utils.UcWorld(path.Base(filename)), -1)
	template = strings.Replace(template, "{:PACKAGE:}", filepath.Base(filepath.Dir(main)), -1)
	utils.CreateFile(main, template)
}

var repositoryTemplate = `package {:PACKAGE:}

type {:REPOSITORY:} struct {}

func (i *{:REPOSITORY:}) {:REPOSITORY:}() string {
	return "" 
}
`