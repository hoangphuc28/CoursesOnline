package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AWS           AWSConfig
	App           AppConfig
	OtherServices OtherServicesConfig
	FireBase      FireBaseConfig
}

type AppConfig struct {
	Version string
	Mode    string
	Port    string
}

type AWSConfig struct {
	Region    string
	APIKey    string
	SecretKey string
	S3Bucket  string
	S3Domain  string
}
type FireBaseConfig struct {
	Bucket              string
	PathCredentialsFile string
}
type OtherServicesConfig struct {
	DurationVideo string
}

func LoadConfig(fileName string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var cf Config
	if err := v.Unmarshal(&cf); err != nil {
		return nil, err
	}
	return &cf, nil
}
