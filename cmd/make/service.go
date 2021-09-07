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
	Service = &cobra.Command{
		Use:                        "make:service",
		Aliases:                    []string{"ms"},
		SuggestFor:                 nil,
		Short:                      "create a service",
		Long:                       "create a service",
		Example:                    "vk make:service demo",
		Run:                        serviceRun,
	}
)

func serviceRun(cmd *cobra.Command, args []string)  {
	curPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if len(args) != 1 {
		cmd.Usage()
		return
	}
	name := args[0]
	log.Println(fmt.Sprintf("create %s service", name))
	generate.CreateService(curPath, path.Join("service", name))
}
