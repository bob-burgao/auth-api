package domain_config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var envVars *Environments

type Environments struct {
	MsName      string        `mapstructure:"MS_NAME"`
	SelfTimeOut time.Duration `mapstructure:"SELF_TIME_OUT"`
	SelfPort    string        `mapstructure:"SELF_PORT"`
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.SetDefault("MS_NAME", "ms-auth")
	viper.SetDefault("SELF_TIME_OUT", 30*time.Second)
	viper.SetDefault("SELF_PORT", "8080")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error read envs")
	}

	if err := viper.Unmarshal(&envVars); err != nil {
		fmt.Println("error unmarshal envs")
	}

	return envVars
}

func EnvVars() *Environments {
	return envVars
}
