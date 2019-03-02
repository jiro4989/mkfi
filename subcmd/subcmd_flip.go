package subcmd

import "github.com/spf13/cobra"

var flipCommand = &cobra.Command{
	Use:   "flip",
	Short: "mkfi",
	Long:  "mkfi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCommand.AddCommand(flipCommand)
	flipCommand.Flags().StringP("owner", "o", "", "owner of remote file")
}
