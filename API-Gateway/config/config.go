package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Services ServicesConfig
	Mysql    MysqlConfig
	Email    EmailConfig
}
type EmailConfig struct {
	AppEmail    string
	AppPassword string
}
type ServicesConfig struct {
	Port                  string
	FileUrl               string
	AuthUrl               string
	CourseUrl             string
	CartUrl               string
	CategoryUrl           string
	PaymentUrl            string
	TopicUrl              string
	Secret                string
	UserUrl               string
	MigrationURL          string
	AccessTokenExpiredIn  int
	RefreshTokenExpiredIn int
}
type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
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
