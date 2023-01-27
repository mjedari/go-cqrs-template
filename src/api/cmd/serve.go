package cmd

import (
	"github.com/mjedari/go-cqrs-template/src/api/config"
	"github.com/mjedari/go-cqrs-template/src/api/route"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving service.",
	Long:  `Serving service.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	go runHttpServer()
}

func runHttpServer() {
	log.WithField("HTTP_Port", ":"+config.Config.Server.Port).
		Info("starting HTTP/REST http...")

	router := route.NewRouter()

	if err := http.ListenAndServe(":"+config.Config.Server.Port, router); err != nil {
		log.Fatal(err)
	}
}
