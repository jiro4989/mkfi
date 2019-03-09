package subcmd

import (
	"bufio"
	"os"

	"github.com/jiro4989/mkfi/domain"
	"github.com/jiro4989/mkfi/usecase"
	"github.com/spf13/cobra"
)

var trimCommand = &cobra.Command{
	Use:   "trim",
	Short: "mkfi",
	Long:  "mkfi",
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flags()

		x, err := f.GetInt("w")
		if err != nil {
			panic(err)
		}
		y, err := f.GetInt("y")
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

		var targets []string
		if 1 < len(args) {
			targets = args
		} else {
			sc := bufio.NewScanner(os.Stdin)
			for sc.Scan() {
				fn := sc.Text()
				targets = append(targets, fn)
			}
		}

		rect := domain.Rectangle{
			X:      x,
			Y:      y,
			Width:  w,
			Height: h,
		}
		usecase.TrimImageFiles(rect, "out/trim", targets)
	},
}

func init() {
	RootCommand.AddCommand(trimCommand)
	trimCommand.Flags().StringP("", "x", "", "Crop X")
	trimCommand.Flags().StringP("", "y", "", "Crop Y")
	trimCommand.Flags().StringP("width", "W", "", "Crop width")
	trimCommand.Flags().StringP("height", "H", "", "Crop height")
}
