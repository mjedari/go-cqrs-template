package cmd

import (
	"fmt"
	"github.com/mjedari/go-cqrs-template/infra/providers/storage"
	"github.com/mjedari/go-cqrs-template/web/config"
	"github.com/mjedari/go-cqrs-template/web/route"
	"github.com/mjedari/go-cqrs-template/web/wiring"

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
	initWiring()
	go runHttpServer()
}

func initWiring() {
	redisProvider, err := storage.NewRedis(config.Config.Redis)
	if err != nil {
		log.Fatalf("Fatal error on create redis connection: %s \n", err)
	}
	wiring.Wiring = wiring.NewWire(redisProvider)
}

func runHttpServer() {
	log.WithField("HTTP_Port", config.Config.Server.Host+":"+config.Config.Server.Port).
		Info("starting HTTP/REST http...")

	router := route.NewRouter()

	address := fmt.Sprintf("%s:%v", config.Config.Server.Host, config.Config.Server.Port)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
