package broker

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "broker",
	Short: "bigworm borker mode",
	Run:   Entry,
}

func Entry(cmd *cobra.Command, args []string) {

}
