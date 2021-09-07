package cmd

import (
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
	rootCmd.Execute()
}