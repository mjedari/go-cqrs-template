package cmd

import (
	"fmt"
	"github.com/mjedari/go-cqrs-template/src/api/config"
	"github.com/mjedari/go-cqrs-template/src/api/wiring"
	"github.com/mjedari/go-cqrs-template/src/infra/providers/storage"
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

	//if configs.Config.Credential.TokenSecret == "" {
	//	log.Fatal("There is no token secret in config file\n")
	//}
	//dbProvider, err := providers.NewPostgresFromConfig(configs.Config.Database)
	//if err != nil {
	//	log.Fatalf("Fatal error on create db: %s \n", err)
	//}
	//cacheProvider, err := providers.NewRedisFromConfig(configs.Config.Cache)
	//if err != nil {
	//	log.Fatalf("Fatal error on create cache connection: %s \n", err)
	//}
	//httpProvider := provider.NewHTTPService(config.Config.HTTPClient)
	redisProvider, err := storage.NewRedis(config.Config.Redis)
	if err != nil {
		log.Fatalf("Fatal error on create redis connection: %s \n", err)
	}
	wiring.Wiring = wiring.NewWire(redisProvider)
	fmt.Println("wire event", wiring.Wiring.GetEventBus())
}
