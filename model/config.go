package model

type AppConfig struct {
	EnvMode  string
	Host     string
	Port     string
	DBdriver string
	DBUrl    string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SSLMode  string
}

