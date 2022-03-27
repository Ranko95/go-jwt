package config

import "os"

type HTTPServerConfig struct {
	Address string
}

type MongoConfig struct {
	Url      string
	User     string
	Password string
	DBName   string
}

type JWTConfig struct {
	Secret string
}

type Config struct {
	HTTPServerConfig HTTPServerConfig
	MongoConfig      MongoConfig
	JWTConfig        JWTConfig
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}

func GetConfig() *Config {
	return &Config{
		HTTPServerConfig: HTTPServerConfig{
			Address: getEnv("HTTP_SERVER_ADDR", "0.0.0.0:8000"),
		},
		MongoConfig: MongoConfig{
			Url:      getEnv("MONGO_URL", ""),
			User:     getEnv("MONGO_USER", ""),
			Password: getEnv("MONGO_PASS", ""),
			DBName:   getEnv("MONGO_DB_NAME", ""),
		},
		JWTConfig: JWTConfig{
			Secret: getEnv("JWT_SECRET", ""),
		},
	}
}
