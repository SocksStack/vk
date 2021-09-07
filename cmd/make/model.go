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
	Model = &cobra.Command{
		Use:                        "make:model",
		SuggestFor:                 nil,
		Short:                      "create a model",
		Long:                       "create a model",
		Example:                    "vk make:model demo",
		Run:                        modelRun,
	}
)

func init() {
	Model.PersistentFlags().String("type", "gorm", "指定 orm 类型, 默认：gorm")
}

func modelRun(cmd *cobra.Command, args []string)  {
	curPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if len(args) != 1 {
		cmd.Usage()
		return
	}
	name := args[0]

	log.Println(fmt.Sprintf("create %s model", name))
	t, _ := cmd.Flags().GetString("type")
	switch t {
	case "gorm":
	default:
		generate.CreateGormModel(curPath, path.Join("model", name))
	}
}
