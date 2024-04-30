package controller

import (
	"log"
	"os"

	"github.com/Samurai1986/auth-service/model"
	"github.com/joho/godotenv"
)

// env variables
const (
	envMode     = "NODE_ENV"
	appHost     = "APPLICATION_HOST"
	appPort     = "APPLICATION_PORT"
	dbDriver    = "DB_DRIVER"
	dbUrl       = "DB_URL"
	envFilename = ".env"
)

// default values
const (
	defaultEnvMode  = "development"
	defaultAppHost  = "0.0.0.0"
	defaultAppPort  = "8000"
	defaultDBDriver = "postgres"
)

var EnvironmentMode = getEnvMode()

// expect only 1 defaultvalue
func getEnv(key string, defaultval ...string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if len(defaultval) == 0 {
		log.Fatalf("Error: ENV value %s not found", key)
	}
	if len(defaultval) > 1 {
		log.Fatalf("find more than 1 default values of %s ENV: %v", key, defaultval)
	}
	log.Printf("Warning: ENV value %s not found. Load default value %s", key, defaultval[0])
	return defaultval[0]
}

// try to load NODE_ENV value. if it not exists loading development environment from .env
func getEnvMode() string {
	value, exists := os.LookupEnv(envMode)
	if !exists {
		err := godotenv.Load(envFilename)
		if err != nil {
			log.Printf("Error loading %s file", envFilename)
		}
		value, exists = os.LookupEnv(envMode)
		if !exists {
			value = "development"
		}
	}
	log.Printf("App starts in %s mode", value)
	return value
}

func NewConfig() *model.AppConfig {
	return &model.AppConfig{
		EnvMode:  EnvironmentMode,
		Host:     getEnv(appHost, defaultAppHost),
		Port:     getEnv(appPort, defaultAppPort),
		DBdriver: getEnv(dbDriver, defaultDBDriver),
		DBUrl:    getEnv(dbUrl),
	}
}
