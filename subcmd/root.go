package subcmd

import "github.com/spf13/cobra"

var RootCommand = &cobra.Command{
	Use:   "mkfi",
	Short: "mkfi",
	Long:  "mkfi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cobra.OnInitialize()
}
