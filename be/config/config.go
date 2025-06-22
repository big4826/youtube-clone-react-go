package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var (
	config *Config
)

type Config struct {
	Environment string       `mapstructure:"environment" validate:"required"`
	Server      ServerConfig `mapstructure:"server" validate:"required"`
	Log         Log          `mapstructure:"log" validate:"required"`
	DB          DBConfig     `mapstructure:"db" validate:"required"`
	RedisConfig RedisConfig  `mapstructure:"redis" validate:"required"`
}

// ServerConfig .
type ServerConfig struct {
	Port         string        `mapstructure:"port" validate:"required"`
	WriteTimeout time.Duration `mapstructure:"writetimeout" validate:"required"`
	ReadTimeout  time.Duration `mapstructure:"readtimeout" validate:"required"`
	IdleTimeout  time.Duration `mapstructure:"idletimeout" validate:"required"`
	CtxTimeout   time.Duration `mapstructure:"ctx_timeout" validate:"required"`
}

// Log .
type Log struct {
	Environment string `mapstructure:"environment" validate:"required"`
	Level       string `mapstructure:"level" validate:"required"`
	Format      string `mapstructure:"format" validate:"required"`
}

// database
type DBConfig struct {
	Host            string        `mapstructure:"host" validate:"required"`
	Port            string        `mapstructure:"port" validate:"required"`
	User            string        `mapstructure:"user" validate:"required"`
	Password        string        `mapstructure:"password" validate:"required"`
	DBname          string        `mapstructure:"dbname" validate:"required"`
	MaxOpenConn     int32         `mapstructure:"maxopenconn" validate:"required"`
	MaxConnLifeTime time.Duration `mapstructure:"maxconnlifetime" validate:"required"`
	MaxIdle         time.Duration `mapstructure:"maxidle" validate:"required"`
}

// redis
type RedisConfig struct {
	Host         string        `mapstructure:"host" validate:"required"`
	Port         string        `mapstructure:"port" validate:"required"`
	Password     string        `mapstructure:"password"`
	DB           int           `mapstructure:"db"`
	PoolTimeout  time.Duration `mapstructure:"pooltimeout" validate:"required"`
	DialTimeout  time.Duration `mapstructure:"dialtimeout" validate:"required"`
	WriteTimeout time.Duration `mapstructure:"writetimeout" validate:"required"`
	ReadTimeout  time.Duration `mapstructure:"readtimeout" validate:"required"`
}

// InitialConfig .
func InitialConfig() *Config {

	configPath, ok := os.LookupEnv("API_CONFIG_PATH")
	if !ok {
		configPath = "./config"
	}

	configName, ok := os.LookupEnv("API_CONFIG_NAME")
	if !ok {
		configName = "config"
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	err := config.validate()
	if err != nil {
		panic(fmt.Sprintf("failed to get configs %s", err.Error()))
	}

	return config
}

func (c *Config) validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
