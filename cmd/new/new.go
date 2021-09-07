package new

import (
	"fmt"
	"github.com/SocksStack/vk/generate"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	New = &cobra.Command{
		Use:                        "new",
		Aliases:                    []string{"n"},
		SuggestFor:                 nil,
		Short:                      "create a new project",
		Long:                       "create a new project",
		Example:                    "vk new demo",
		Run:                        run,
	}
)

func init() {
}

func run(cmd *cobra.Command, args []string)  {
	curPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if len(args) != 1 {
		cmd.Usage()
		return
	}
	moduleName := args[0]
	curPath = path.Join(curPath, moduleName)
	log.Println(fmt.Sprintf("create %s module", moduleName))
	log.Println(fmt.Sprintf("go mod init %s", moduleName))
	if err := generate.Create(moduleName, curPath); err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("create main.go")
	generate.CreateMain(moduleName, curPath, "main")

	log.Println("create README.md")
	generate.CreateReadMe(moduleName, curPath)

	log.Println("create config")
	generate.CreateConfig(filepath.Join(curPath, "config"))

	log.Println("create .gitignore")
	generate.CreateGitignore(moduleName, curPath)

	log.Println("create bootstrap/bootstrap.go")
	generate.CreateBootstrap(moduleName, curPath, "bootstrap/bootstrap")

	log.Println("create middleware/init.go")
	generate.CreateMiddlewareInit(curPath, "middleware/init")

	log.Println("create router/router.go")
	generate.CreateRouter(moduleName, curPath, "router/router")

	log.Println("create handler/demo.go")
	generate.CreateHandler(curPath, "handler/index")

	generate.ModTidy(moduleName)
}

