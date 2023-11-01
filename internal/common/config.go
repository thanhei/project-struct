package common

import "github.com/spf13/viper"

type serviceConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type mysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type systemConfig struct {
	Secret      string `mapstructure:"secret"`
	TokenExpire int    `mapstructure:"token_expire"`
}

type Config struct {
	Service serviceConfig `json:"service" mapstructure:"service"`
	Mysql   mysqlConfig   `json:"mysql" mapstructure:"mysql"`
	System  systemConfig  `json:"system" mapstructure:"system"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
