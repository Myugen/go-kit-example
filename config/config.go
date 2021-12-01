package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	Server struct {
		Port string
	}
}

var Config = &Configuration{}

type ConfigFlag struct {
	Field string
	Flag  *pflag.Flag
}

func Init(flags ...ConfigFlag) func() {
	return func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		viper.AddConfigPath("/etc/app/")
		viper.AddConfigPath("$HOME/.app")
		viper.AddConfigPath(".")

		viper.SetEnvPrefix("APP")
		viper.AutomaticEnv()

		for _, flag := range flags {
			if err := viper.BindPFlag(flag.Field, flag.Flag); err != nil {
				log.Fatalf("error binding command %s flag: %s", flag.Field, err)
			}
		}

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("error reading configuration: %s", err)
		}

		if err := viper.Unmarshal(Config); err != nil {
			log.Fatalf("error setting confinguration: %s", err)
		}
	}
}
