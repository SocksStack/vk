package cmd

import (
	"fmt"
	m "github.com/SocksStack/vk/cmd/make"
	n "github.com/SocksStack/vk/cmd/new"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{}
)

func init() {
	rootCmd.AddCommand(
		n.New,
		m.Service,
		m.Middleware,
		m.Handler,
		m.Repository,
		m.Logic,
	)
}

func Execute()  {
	fmt.Println(banner)
	rootCmd.Execute()
}

var banner = `
 __   __   __  __       
/\ \ / /  /\ \/ /       
\ \ \'/   \ \  _"-.     
 \ \__|    \ \_\ \_\    
  \/_/      \/_/\/_/    
   enjoy to use!!!
`