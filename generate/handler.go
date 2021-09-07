package generate

import (
	"fmt"
	"github.com/SocksStack/vk/utils"
	"path"
	"path/filepath"
	"strings"
)

func CreateHandler(curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	template := strings.Replace(handlerTemplate, "{:HANDLER:}", utils.UcWorld(filepath.Base(filename)), -1)
	template = strings.Replace(template, "{:PACKAGE:}", filepath.Base(filepath.Dir(main)), -1)
	utils.CreateFile(main, template)
}

var handlerTemplate = `package {:PACKAGE:}

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type {:HANDLER:} struct {}

func (handler *{:HANDLER:}) index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "hello world"})
}

func (handler *{:HANDLER:}) Route(engine *gin.Engine) {
	engine.GET("/{:HANDLER:}", handler.index)
}
`
