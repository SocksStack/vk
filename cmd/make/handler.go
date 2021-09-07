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
	Handler = &cobra.Command{
		Use:                        "make:handler",
		Aliases:                    []string{"mh"},
		SuggestFor:                 nil,
		Short:                      "create a handler",
		Long:                       "create a handler",
		Example:                    "vk make:handler demo or vk mh demo",
		Run:                        handlerRun,
	}
)

func handlerRun(cmd *cobra.Command, args []string)  {
	curPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if len(args) != 1 {
		cmd.Usage()
		return
	}
	name := args[0]
	log.Println(fmt.Sprintf("create %s handler", name))
	generate.CreateHandler(curPath, path.Join("handler", name))
}
