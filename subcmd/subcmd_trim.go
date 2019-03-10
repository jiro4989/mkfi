package subcmd

import (
	"bufio"
	"os"

	"github.com/jiro4989/mkfi/domain"
	"github.com/jiro4989/mkfi/log"
	"github.com/jiro4989/mkfi/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(trimCommand)
	trimCommand.Flags().IntP("axis-x", "x", 0, "Crop X")
	trimCommand.Flags().IntP("axis-y", "y", 0, "Crop Y")
	trimCommand.Flags().IntP("width", "", 0, "Crop width")
	trimCommand.Flags().IntP("height", "", 0, "Crop height")
}

var trimCommand = &cobra.Command{
	Use:   "trim",
	Short: "mkfi",
	Long:  "mkfi",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("start 'trim' subcommand.")
		f := cmd.Flags()

		log.Debug("get commandline option parameters.")
		x, err := f.GetInt("axis-x")
		if err != nil {
			panic(err)
		}
		y, err := f.GetInt("axis-y")
		if err != nil {
			panic(err)
		}
		w, err := f.GetInt("width")
		if err != nil {
			panic(err)
		}
		h, err := f.GetInt("height")
		if err != nil {
			panic(err)
		}
		log.Debug("parameters:x=", x, ",y=", y, ",width=", w, ",height=", h)

		var targets []string
		if 1 <= len(args) {
			log.Debug("target input is arguments.")
			targets = args
		} else {
			log.Debug("target input is stdin.")
			sc := bufio.NewScanner(os.Stdin)
			for sc.Scan() {
				fn := sc.Text()
				targets = append(targets, fn)
			}
			if err := sc.Err(); err != nil {
				panic(err)
			}
		}
		log.Debug("target inputs are ", targets, ".")

		rect := domain.Rectangle{
			X:      x,
			Y:      y,
			Width:  w,
			Height: h,
		}
		usecase.TrimImageFiles(rect, "out/trim", targets)
		log.Debug("end 'trim' subcommand")
	},
}
