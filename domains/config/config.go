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

	DynamoUrl string `mapstructure:"DYNAMO_URL"`
	AwsRegion string `mapstructure:"AWS_REGION"`
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.SetDefault("MS_NAME", "ms-auth")
	viper.SetDefault("SELF_TIME_OUT", 30*time.Second)
	viper.SetDefault("SELF_PORT", "8080")

	viper.SetDefault("DYNAMO_URL", "https://localhost.localstack.cloud:4566")
	viper.SetDefault("AWS_REGION", "us-east-1")

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
