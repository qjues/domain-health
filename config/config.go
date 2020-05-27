package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

var Instance *Config

type Config struct {
	DebugLevel string `mapstructure:"debug_level"`
	StoreDir   string `mapstructure:"store_dir"`
	LogFile    string `mapstructure:"log_file"`

	ListenAddress string `mapstructure:"listen_address"`
	HealthTime    int    `mapstructure:"health_time"`
}

func init() {
	getwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath("/etc/domain-health")
	viper.AddConfigPath(getwd)
	viper.AddConfigPath(executable)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yaml")
	viper.SetConfigName("config.local.yaml")

	viper.SetEnvPrefix("dh")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	Instance = &Config{}
	err = viper.Unmarshal(Instance)

	if err != nil {
		panic(err)
	}
}
