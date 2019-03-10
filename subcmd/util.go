package subcmd

import (
	"bufio"
	"os"

	"github.com/jiro4989/mkfi/log"
)

func fetchTargetFiles(args []string) []string {
	if 1 <= len(args) {
		log.Debug("target input is arguments. targets=", args)
		return args
	}

	log.Debug("target input is stdin.")
	var targets []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fn := sc.Text()
		targets = append(targets, fn)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	log.Debug("target input is stdin. targets=", targets)
	return targets
}
