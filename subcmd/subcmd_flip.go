package subcmd

import (
	"github.com/jiro4989/mkfi/log"
	"github.com/jiro4989/mkfi/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(flipCommand)
	flipCommand.Flags().StringP("outdir", "o", "out/trim", "Save dir")
}

var flipCommand = &cobra.Command{
	Use:   "flip",
	Short: "mkfi",
	Long:  "mkfi",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("start 'flip' subcommand.")

		f := cmd.Flags()

		log.Debug("get commandline option parameters.")
		outDir, err := f.GetString("outdir")
		if err != nil {
			panic(err)
		}
		log.Debug("commandline options:outDir=", outDir)

		targets := fetchTargetFiles(args)
		usecase.FlipImageFiles(outDir, targets)
		log.Debug("end 'flip' subcommand.")
	},
}
