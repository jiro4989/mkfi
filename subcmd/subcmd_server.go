package subcmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jiro4989/mkfi/api"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "mkfi",
	Long:  "mkfi",
	Run: func(cmd *cobra.Command, args []string) {
		router := httprouter.New() // HTTPルーターを初期化

		router.GET("/", api.RootPage)
		router.POST("/generate-chain", api.GenerateChain)
		router.POST("/save", api.Save)
		router.POST("/generate", api.Generate)
		router.POST("/trim", api.Trim)
		router.POST("/flip", api.Flip)
		router.POST("/paste", api.Paste)

		const port = "8080"

		// Webサーバーを8080ポートで立ち上げる
		fmt.Println("http://localhost:" + port + "/")
		err := http.ListenAndServe(":"+port, router)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCommand.AddCommand(serverCommand)
	serverCommand.Flags().StringP("owner", "o", "", "owner of remote file")
}
