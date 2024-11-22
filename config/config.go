package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var config Config

type Config struct {
	ServerConfig   ServerConfig
	PostgresConfig PostgresConfig
	AWSConfig      AWSConfig
}

type ServerConfig struct {
	Port string
}

type PostgresConfig struct {
	Host     string `envconfig:"POSTGRES_HOST"`
	Port     string `envconfig:"POSTGRES_PORT"`
	User     string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	DBName   string `envconfig:"POSTGRES_DB"`
}

type AWSConfig struct {
	Region          string `envconfig:"AWS_REGION"`
	AccessKeyID     string `envconfig:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	BucketName      string `envconfig:"AWS_BUCKET_NAME"`
}

func GetConfig() *Config {
	return &config
}

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return envconfig.Process("", &config)
}
