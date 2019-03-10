package subcmd

import "github.com/spf13/cobra"

func init() {
	RootCommand.AddCommand(generateCommand)
	generateCommand.Flags().StringP("owner", "o", "", "owner of remote file")
}

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "mkfi",
	Long:  "mkfi",
	Run:   func(cmd *cobra.Command, args []string) {},
}
