package generate

import (
	"fmt"
	"github.com/SocksStack/vk/utils"
	"path"
	"path/filepath"
	"strings"
)

func CreateMiddleware(curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	template := strings.Replace(middlewareTemplate, "{:MIDDLEWARE:}", utils.UcWorld(path.Base(filename)), -1)
	template = strings.Replace(template, "{:PACKAGE:}", filepath.Base(filepath.Dir(main)), -1)
	utils.CreateFile(main, template)
}

var middlewareTemplate = `package {:PACKAGE:}

import "github.com/gin-gonic/gin"

func {:MIDDLEWARE:}() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}
`

func CreateMiddlewareInit(curPath, filename string)  {
	main := path.Join(curPath, fmt.Sprintf("%s.go", filename))
	utils.CreateFile(main, middlewareInitTemplate)
}

var middlewareInitTemplate = `package middleware

import "github.com/SocksStack/common/constract"

func Middlewares() constract.Middlewares {
	return constract.Middlewares{
	
	}
}
`