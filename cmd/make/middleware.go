package make

import (
	"fmt"
	"github.com/SocksStack/vk/generate"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

var (
	Middleware = &cobra.Command{
		Use:                        "make:middleware",
		Aliases:                    []string{"mm"},
		SuggestFor:                 nil,
		Short:                      "create a middleware",
		Long:                       "create a middleware",
		Example:                    "vk make:middleware demo or vk mm demo",
		Run:                        middlewareRun,
	}
)

func middlewareRun(cmd *cobra.Command, args []string)  {
	curPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if len(args) != 1 {
		cmd.Usage()
		return
	}
	name := args[0]
	log.Println(fmt.Sprintf("create %s middleware", name))
	generate.CreateMiddleware(curPath, path.Join("middleware", name))
}
