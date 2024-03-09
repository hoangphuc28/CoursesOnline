package config

import "github.com/spf13/viper"

type Config struct {
	Service       ServiceConfig
	Mysql         mysqlConfig
	Email         EmailConfig
	OtherServices OtherServices
}
type ServiceConfig struct {
	Version               string
	Mode                  string
	Port                  string
	Secret                string
	PasswordTokenExpired  int
	AccessTokenExpiredIn  int
	RefreshTokenExpiredIn int
	MaxSizeMess           int
}
type EmailConfig struct {
	AppEmail    string
	AppPassword string
}
type mysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}
type OtherServices struct {
	FileUrl        string
	CartServiceUrl string
	PaymentUrl     string
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
