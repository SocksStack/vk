package generate

import (
	"fmt"
	"github.com/SocksStack/vk/utils"
	"path"
	"path/filepath"
	"strings"
)

func CreateGormModel(curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	template := strings.Replace(gormModelTemplate, "{:MODEL:}", utils.UcWorld(path.Base(filename)), -1)
	template = strings.Replace(template, "{:SNAKEMODEL:}", utils.SnakeString(path.Base(filename)), -1)
	template = strings.Replace(template, "{:PACKAGE:}", filepath.Base(filepath.Dir(main)), -1)
	utils.CreateFile(main, template)
}


var gormModelTemplate = `package {:PACKAGE:}

import "gorm.io/gorm"

type {:MODEL:} struct {
	gorm.Model
}

func (m {:MODEL:}) TableName() string {
	return "{:SNAKEMODEL:}"
}`
