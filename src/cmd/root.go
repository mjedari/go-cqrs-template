package cmd

import (
	"github.com/mjedari/go-cqrs-template/web/config"
	log "github.com/sirupsen/logrus"
)
import "github.com/spf13/cobra"
import "github.com/spf13/viper"

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:   "my-cqrs-template",
		Short: "short description",
		Long:  `long description`,
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Mahdi Jedari", "i.jedari@gmail.com")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	viper.Unmarshal(&config.Config)
	log.Info("configuration initialized!")
}
