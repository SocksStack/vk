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
	Repository = &cobra.Command{
		Use:                        "make:repository",
		Aliases:                    []string{"mr"},
		SuggestFor:                 nil,
		Short:                      "create a repository",
		Long:                       "create a repository",
		Example:                    "vk make:repository demo or vk mr demo",
		Run:                        repositoryRun,
	}
)

func repositoryRun(cmd *cobra.Command, args []string)  {
	curPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if len(args) != 1 {
		cmd.Usage()
		return
	}
	name := args[0]
	log.Println(fmt.Sprintf("create %s repository", name))
	generate.CreateRepository(curPath, path.Join("repository", name))
}
