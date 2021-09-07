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
	Logic = &cobra.Command{
		Use:                        "make:logic",
		Aliases:                    []string{"ml"},
		SuggestFor:                 nil,
		Short:                      "create a logic",
		Long:                       "create a logic",
		Example:                    "vk make:logic demo or vk ml demo",
		Run:                        logicRun,
	}
)

func logicRun(cmd *cobra.Command, args []string)  {
	curPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if len(args) != 1 {
		cmd.Usage()
		return
	}
	name := args[0]
	log.Println(fmt.Sprintf("create %s logic", name))
	generate.CreateLogic(curPath, path.Join("logic", name))
}

