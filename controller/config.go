package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/Samurai1986/auth-service/model"
	"github.com/joho/godotenv"
)

// env variables
var (
	envMode     = "NODE_ENV"
	appHost     = "APPLICATION_HOST"
	appPort     = "APPLICATION_PORT"
	dbDriver    = "DB_DRIVER"
	dbHost      = "DB_HOST"
	dbPort      = "DB_PORT"
	dbUser      = "DB_USER"
	dbPass      = "DB_PASSWORD"
	dbName      = "DB_NAME"
	dbSSLMode   = "DB_SSL_MODE"
	envFilename = "../.env"
)

func GetPostgresDBUrl(p *model.PostgresConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.Database,
		p.SSLMode,
	)
}

func getEnv(key string, defaultval string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultval
}

// Todo: more simple
func NewConfig() *model.AppConfig {
	var config *model.AppConfig
	appEnvMode := getEnv(envMode, "development")
	if appEnvMode == "production" {
		err := godotenv.Load(envFilename)
		if err != nil {
			log.Fatal("Error loading.env file")
		}
		config = &model.AppConfig{
			EnvMode:  appEnvMode,
			Host:     os.Getenv(appHost),
			Port:     os.Getenv(appPort),
			DBdriver: os.Getenv(dbDriver),
			DBUrl: GetPostgresDBUrl(&model.PostgresConfig{
				Host:     os.Getenv(dbHost),
				Port:     os.Getenv(dbPort),
				User:     os.Getenv(dbUser),
				Database: os.Getenv(dbName),
				Password: os.Getenv(dbPass),
				SSLMode:  os.Getenv(dbSSLMode),
			}),
		}
		return config
	}
	config = loadDevEnv()
	return config
}

// in dev mode app get values from .env file
// in dev mode if value is not configured it has a default value
func loadDevEnv() *model.AppConfig {
	_ = godotenv.Load(envFilename)
	return &model.AppConfig{
		EnvMode:  getEnv(envMode, "development"),
		Host:     getEnv(appHost, "0.0.0.0"),
		Port:     getEnv(appPort, "8000"),
		DBdriver: getEnv(dbDriver, "postgres"),
		DBUrl: GetPostgresDBUrl(&model.PostgresConfig{
			User:     getEnv(dbUser, "postgres"),
			Password: getEnv(dbPass, "postgres"),
			Host:     getEnv(dbHost, "localhost"),
			Port:     getEnv(dbPort, "5432"),
			Database: getEnv(dbName, "postgres"),
			SSLMode:  getEnv(dbSSLMode, "false"),
		}),
	}
}
