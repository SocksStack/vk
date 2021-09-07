package generate

import (
	"github.com/SocksStack/vk/utils"
	"path"
)

func CreateGitignore(moduleName, curPath string)  {
	main := path.Join(curPath, ".gitignore")
	utils.CreateFile(main, "")
}
