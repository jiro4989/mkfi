package subcmd

import (
	"github.com/jiro4989/mkfi/log"
	"github.com/jiro4989/mkfi/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(scaleCommand)
	scaleCommand.Flags().IntP("scalesize", "s", 100, "scale size")
	scaleCommand.Flags().StringP("outdir", "d", "out/scale", "out dir")
}

var scaleCommand = &cobra.Command{
	Use:   "scale",
	Short: "mkfi",
	Long:  "mkfi",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("start 'scale' subcommand.")
		f := cmd.Flags()

		log.Debug("get commandline option parameters.")
		scaleSize, err := f.GetInt("scalesize")
		if err != nil {
			panic(err)
		}
		outDir, err := f.GetString("outdir")
		if err != nil {
			panic(err)
		}
		log.Debug("commandline options:scalesize=", scaleSize, ",outdir=", outDir)

		targets := fetchTargetFiles(args)
		usecase.ScaleImageFiles(scaleSize, outDir, targets)
		log.Debug("end 'scale' subcommand")
	},
}
