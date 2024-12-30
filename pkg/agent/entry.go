package agent

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "agent",
	Short: "bigworm agent mode",
	Run:   Entry,
}

func Entry(cmd *cobra.Command, args []string) {

}
