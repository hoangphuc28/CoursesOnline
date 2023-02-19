package config

import "github.com/spf13/viper"

type Config struct {
	Service       ServiceConfig
	Mysql         mysqlConfig
	Email         EmailConfig
	OtherServices otherServicesConfig
	ClientSide    ClientSideConfig
	RabbitMq      RabbitMq
}
type RabbitMq struct {
	User     string
	Password string
	Port     string
	Host     string
}
type ServiceConfig struct {
	Version               string
	Mode                  string
	Port                  string
	Secret                string
	ActiveTokenExpired    int
	AccessTokenExpiredIn  int
	RefreshTokenExpiredIn int
}
type mysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}
type ClientSideConfig struct {
	URL string
}
type EmailConfig struct {
	AppEmail    string
	AppPassword string
}
type otherServicesConfig struct {
	MailServiceURL string
	CartServiceUrl string
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
