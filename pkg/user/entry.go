package user

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "user ",
	Short: "bigworm user mode",
	Run:   Entry,
}

func Entry(cmd *cobra.Command, args []string) {

}
