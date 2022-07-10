package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Configure maps configuration files to structs
func Configure(conf interface{}, fileName string) {
	viper.SetConfigName(fileName)
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file read error")
		fmt.Println(err)
		os.Exit(1)
	}

	// UnmarshalしてCにマッピング
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}
