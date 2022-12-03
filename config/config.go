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

var _ config = (*configData)(nil)

type config interface {
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
	RedisHost() string
	RedisPort() string
	RedisDB() int
	CacheTimeout() time.Duration
	CacheExpiration() time.Duration
	Pagination() int
}

//nolint:revive
type configData struct {
	AppName_          string `mapstructure:"APP_NAME"`
	APIAddress_       string `mapstructure:"API_ADDRESS"`
	IdleTimeout_      int    `mapstructure:"IDLE_TIMEOUT"`
	ReadTimeout_      int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout_     int    `mapstructure:"WRITE_TIMEOUT"`
	PostgresHost_     string `mapstructure:"PG_HOST"`
	PostgresPort_     string `mapstructure:"PG_PORT"`
	PostgresName_     string `mapstructure:"PG_NAME"`
	PostgresUser_     string `mapstructure:"PG_USER"`
	PostgresPassword_ string `mapstructure:"PG_PASSWORD"`
	RedisHost_        string `mapstructure:"REDIS_HOST"`
	RedisPort_        string `mapstructure:"REDIS_PORT"`
	RedisDB_          int    `mapstructure:"REDIS_DB"`
	CacheTimeout_     int    `mapstructure:"CACHE_TIMEOUT"`
	CacheExpiration_  int    `mapstructure:"CACHE_EXPIRATION"`
	Pagination_       int    `mapstructure:"PAGINATION"`
}

func New() (*configData, error) {
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
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("CACHE_TIMEOUT", 100)
	viper.SetDefault("CACHE_EXPIRATION", 25)
	viper.SetDefault("PAGINATION", 10)
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

func (cfg *configData) RedisHost() string {
	return cfg.RedisHost_
}

func (cfg *configData) RedisPort() string {
	return cfg.RedisPort_
}

func (cfg *configData) RedisDB() int {
	return cfg.RedisDB_
}

func (cfg *configData) CacheTimeout() time.Duration {
	return time.Duration(cfg.CacheTimeout_) * time.Millisecond
}

func (cfg *configData) CacheExpiration() time.Duration {
	return time.Duration(cfg.CacheExpiration_) * time.Second
}

func (cfg *configData) Pagination() int {
	return cfg.Pagination_
}
