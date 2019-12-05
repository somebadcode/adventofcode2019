package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var config *viper.Viper

func main() {
	logger := log.New(os.Stdout, "", 0)
	config = viper.New()

	config.SetDefault("testdata", "testdata")
	config.SetDefault("inputdirectory", "inputdirectory")

	config.AddConfigPath("./configs")
	config.AddConfigPath(".")
	config.SetConfigName("config")
	config.AutomaticEnv()
	err := config.ReadInConfig()
	if err != nil {
		logger.Fatalln(err)
	}

	err = solve(config.GetString("inputdirectory"), config, logger)
	if err != nil {
		os.Exit(1)
	}
}
