package subcmd

import "github.com/spf13/cobra"

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "mkfi",
	Long:  "mkfi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCommand.AddCommand(serverCommand)
	serverCommand.Flags().StringP("owner", "o", "", "owner of remote file")
}
