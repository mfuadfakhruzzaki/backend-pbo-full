// config/config.go
package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	JWTSecret          string
	JWTExpiration      time.Duration
	CORSAllowedOrigins []string
	RateLimit          int
	RateLimitWindow    time.Duration
	GoogleClientID     string
	GoogleClientSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading configuration from environment variables")
	}

	rateLimit, err := strconv.Atoi(getEnv("RATE_LIMIT", "100"))
	if err != nil {
		rateLimit = 100 // default value
	}

	rateLimitWindow, err := time.ParseDuration(getEnv("RATE_LIMIT_WINDOW", "15m"))
	if err != nil {
		rateLimitWindow = 15 * time.Minute
	}

	jwtExpiration, err := time.ParseDuration(getEnv("JWT_EXPIRATION", "24h"))
	if err != nil {
		jwtExpiration = 24 * time.Hour
	}

	corsOrigins := getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	allowedOrigins := splitAndTrim(corsOrigins, ",")

	return &Config{
		Port:               getEnv("PORT", "8080"),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "your_db_user"),
		DBPassword:         getEnv("DB_PASSWORD", "your_db_password"),
		DBName:             getEnv("DB_NAME", "dashboard_db"),
		JWTSecret:          getEnv("JWT_SECRET", "your_jwt_secret_key"),
		JWTExpiration:      jwtExpiration,
		CORSAllowedOrigins: allowedOrigins,
		RateLimit:          rateLimit,
		RateLimitWindow:    rateLimitWindow,
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func splitAndTrim(s, sep string) []string {
	var result []string
	for _, part := range strings.Split(s, sep) {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
