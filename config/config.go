package config

import (
    "github.com/spf13/viper"
    "log"
)

func LoadConfig() {
    viper.SetConfigFile(".env")
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
    viper.AutomaticEnv()
}