package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config has all configuration for the project
type Config struct {
	ServerConfig ServerConfig `mapstructure:"server"`
	CORSConfig   CORSConfig   `mapstructure:"cors"`
	LoggerConfig LoggerConfig `mapstructure:"logger"`
}

// ServerConfig has only server specific configuration
type ServerConfig struct {
	Port         string        `mapstructure:"port"`
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	CloseTimeout time.Duration `mapstructure:"closeTimeout"`
}

// CORSConfig contains CORS related configurations
type CORSConfig struct {
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
	AllowedMethods   []string `mapstructure:"allowed_methods"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	AllowedHeaders   []string `mapstructure:"allowed_headers"`
}

// LoggerConfig has logger related configuration.
type LoggerConfig struct {
	LogFilePath string `mapstructure:"file"`
}

// MiddlewareConfig has middlewares related configuration
type MiddlewareConfig struct {
	EnableRequestLog bool `mapstructure:"enableRequestLog"`
}

// GetConfig returns entire project configuration
func GetConfig() *Config {
	return GetConfigFromFile("default")
}

// GetConfigFromFile returns configuration from specific file object
func GetConfigFromFile(fileName string) *Config {
	if fileName == "" {
		fileName = "default"
	}

	// looking for filename `default` inside `src/server` dir with `.toml` extension
	viper.SetConfigName(fileName)
	viper.AddConfigPath("../conf/")
	viper.AddConfigPath("../../conf/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("couldn't load config: %s", err)
		os.Exit(1)
	}
	config := &Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("couldn't read config: %s", err)
		os.Exit(1)
	}
	return config
}
