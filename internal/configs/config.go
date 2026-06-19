package configs

import (
	"backendServer/internal/consts"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".system.env")
	if err != nil {
		log.Fatal("Error loading .system.env file: ", err)
	}
}

type AppConfig struct {
	Port          string
	GinMode       string
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
	LogCompress   bool
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		Port:          getEnvAsStr("PORT", consts.DefaultPort),
		GinMode:       getEnvAsStr("GIN_MODE", consts.DefaultGinMode),
		LogMaxSize:    getEnvAsInt("LOG_MAX_SIZE", 100),
		LogMaxBackups: getEnvAsInt("LOG_MAX_BACKUPS", 3),
		LogMaxAge:     getEnvAsInt("LOG_MAX_AGE", 28),
		LogCompress:   getEnvAsBool("LOG_COMPRESS", true),
	}
}

func getEnvAsStr(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	return val
}

func getEnvAsInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return intVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}
	return boolVal
}
