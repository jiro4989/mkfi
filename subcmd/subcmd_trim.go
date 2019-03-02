package subcmd

import "github.com/spf13/cobra"

var trimCommand = &cobra.Command{
	Use:   "trim",
	Short: "mkfi",
	Long:  "mkfi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCommand.AddCommand(trimCommand)
	trimCommand.Flags().StringP("owner", "o", "", "owner of remote file")
}
