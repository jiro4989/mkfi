package subcmd

import "github.com/spf13/cobra"

func init() {
	RootCommand.AddCommand(pasteCommand)
	pasteCommand.Flags().StringP("owner", "o", "", "owner of remote file")
}

var pasteCommand = &cobra.Command{
	Use:   "paste",
	Short: "mkfi",
	Long:  "mkfi",
	Run:   func(cmd *cobra.Command, args []string) {},
}
