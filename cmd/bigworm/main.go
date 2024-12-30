package main

import (
	"bigworm/pkg/agent"
	"bigworm/pkg/broker"
	"bigworm/pkg/user"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bigworm",
	Short: "Hugo is a very fast static site generator",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(agent.RootCmd)
	rootCmd.AddCommand(broker.RootCmd)
	rootCmd.AddCommand(user.RootCmd)
}

func main() {

}
