package subcmd

import "github.com/jiro4989/mkfi/global"
import "github.com/spf13/cobra"

var RootCommand = &cobra.Command{
	Use:   "mkfi",
	Short: "mkfi",
	Long:  "mkfi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cobra.OnInitialize()
	RootCommand.PersistentFlags().BoolVarP(&global.DebugFlag, "debug", "X", false, "Debug logging flag.")
}
