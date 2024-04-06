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
	envFilename = ".env"
)

var EnvironmentMode = getEnvMode()



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
	log.Printf("ENV value %s not found. Load default value %s", key, defaultval)
	return defaultval
}

//load .env file and try to load NODE_ENV value. if it not exists loading 
func getEnvMode() string {
	err := godotenv.Load(envFilename)
	if err != nil {
		log.Printf("Error loading %s file", envFilename)
	}
	value := getEnv(envMode, "development")
	log.Printf("App starts in %s mode", value)
	return value
}

func NewConfig() *model.AppConfig {
	return &model.AppConfig{
		EnvMode:  EnvironmentMode,
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
