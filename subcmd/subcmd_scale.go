package subcmd

import "github.com/spf13/cobra"

func init() {
	RootCommand.AddCommand(scaleCommand)
	scaleCommand.Flags().StringP("owner", "o", "", "owner of remote file")
}

var scaleCommand = &cobra.Command{
	Use:   "scale",
	Short: "mkfi",
	Long:  "mkfi",
	Run:   func(cmd *cobra.Command, args []string) {},
}
