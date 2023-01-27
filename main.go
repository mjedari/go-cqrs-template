package main

import (
	"github.com/mjedari/go-cqrs-template/src/api/cmd"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {

	cmd.Execute()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	// Perform any necessary cleanup before exiting
	log.Info("Exiting...")
}
