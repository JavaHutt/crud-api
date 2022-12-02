package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	configFileName = "app"
	configFileType = "yaml"
)

type Config interface {
	AppName() string
	APIAddress() string
	IdleTimeout() time.Duration
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	PostgresHost() string
	PostgresPort() string
	PostgresName() string
	PostgresUser() string
	PostgresPassword() string
}

type configData struct {
	AppName_          string        `mapstructure:"APP_NAME"`
	APIAddress_       string        `mapstructure:"API_ADDRESS"`
	IdleTimeout_      time.Duration `mapstructure:"IDLE_TIMEOUT"`
	ReadTimeout_      time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout_     time.Duration `mapstructure:"WRITE_TIMEOUT"`
	PostgresHost_     string        `mapstructure:"PG_HOST"`
	PostgresPort_     string        `mapstructure:"PG_PORT"`
	PostgresName_     string        `mapstructure:"PG_NAME"`
	PostgresUser_     string        `mapstructure:"PG_USER"`
	PostgresPassword_ string        `mapstructure:"PG_PASSWORD"`
}

func New() (Config, error) {
	cfg, err := configureViper()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func configureViper() (*configData, error) {
	setDefaults()

	var err error
	if err = loadConfigToViper("."); err != nil {
		fmt.Printf("cannot load the config file: %v\n", err)
	}

	viper.AutomaticEnv()

	var cfg configData
	err = viper.Unmarshal(&cfg)
	return &cfg, err
}

func setDefaults() {
	viper.SetDefault("APP_NAME", "crud-api")
	viper.SetDefault("API_ADDRESS", ":3000")
	viper.SetDefault("IDLE_TIMEOUT", 1)
	viper.SetDefault("READ_TIMEOUT", 1)
	viper.SetDefault("WRITE_TIMEOUT", 1)
	viper.SetDefault("POSTGRES_HOST", "localhost")
	viper.SetDefault("POSTGRES_PORT", "5432")
	viper.SetDefault("POSTGRES_NAME", "crud")
	viper.SetDefault("POSTGRES_USER", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "postgres")
}

func loadConfigToViper(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileType)
	return viper.ReadInConfig()
}

func (cfg *configData) AppName() string {
	return cfg.AppName_
}

func (cfg *configData) APIAddress() string {
	return cfg.APIAddress_
}

func (cfg *configData) IdleTimeout() time.Duration {
	return time.Duration(cfg.IdleTimeout_) * time.Minute
}

func (cfg *configData) ReadTimeout() time.Duration {
	return time.Duration(cfg.ReadTimeout_) * time.Minute
}

func (cfg *configData) WriteTimeout() time.Duration {
	return time.Duration(cfg.WriteTimeout_) * time.Minute
}

func (cfg *configData) PostgresHost() string {
	return cfg.PostgresHost_
}

func (cfg *configData) PostgresPort() string {
	return cfg.PostgresPort_
}

func (cfg *configData) PostgresName() string {
	return cfg.PostgresName_
}

func (cfg *configData) PostgresUser() string {
	return cfg.PostgresUser_
}

func (cfg *configData) PostgresPassword() string {
	return cfg.PostgresPassword_
}
