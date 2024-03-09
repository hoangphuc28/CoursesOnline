package config

import "github.com/spf13/viper"

type Config struct {
	Service       ServiceConfig
	Mysql         mysqlConfig
	Paypal        paypalConfig
	ClientSide    clientSide
	OtherServices OtherServices
}

type ServiceConfig struct {
	Version string
	Mode    string
	Port    string
	Secret  string

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
type paypalConfig struct {
	ClientId        string
	SecretKey       string
	BaseUrl         string
	CreateOrderApi  string
	CaptureOrderApi string
	IdentifyApi     string
	PayoutApi       string
	GetAccessToken  string
}
type clientSide struct {
	Url string
}
type OtherServices struct {
	CartServiceUrl   string
	CourseServiceUrl string
}

func LoadConfig(fileName string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(fileName)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var cf Config
	if err := v.Unmarshal(&cf); err != nil {
		return nil, err
	}

	return &cf, nil
}
