package config

import (

	"github.com/spf13/viper"
)

func LoadConfig() {
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatalf("Error reading config file: %v", err)
	// }
	viper.AutomaticEnv()
}
