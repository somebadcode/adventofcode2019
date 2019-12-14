package main

import "github.com/spf13/viper"

func mustGetConfig() *viper.Viper {
	v, err := getConfig()
	if err != nil {
		panic(err)
	}
	return v
}

func getConfig() (*viper.Viper, error) {
	config = viper.New()

	config.SetDefault("testdata", "testdata")
	config.SetDefault("inputdirectory", "inputdirectory")

	config.AddConfigPath("./configs")
	config.AddConfigPath(".")
	config.SetConfigName("config")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	return config, err
}
