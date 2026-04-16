package config

import (
	"os"
	"strconv"
)

// Default configuration values
const (
	DefaultPort        = 2095
	DefaultWebBasePath = "/"
	DefaultDBPath      = "db/s-ui.db"
	DefaultLogLevel    = "info"
	DefaultLogFile     = ""
	DefaultLogMaxSize  = 10 // MB
)

// Config holds the application configuration
type Config struct {
	// Server settings
	Port        int
	WebBasePath string

	// Database settings
	DBPath string

	// Logging settings
	LogLevel   string
	LogFile    string
	LogMaxSize int

	// Security settings
	SecretKey string
}

// GetConfig returns the application configuration, reading from environment
// variables with fallback to default values.
func GetConfig() *Config {
	return &Config{
		Port:        getEnvInt("SUI_PORT", DefaultPort),
		WebBasePath: getEnvStr("SUI_WEB_BASE_PATH", DefaultWebBasePath),
		DBPath:      getEnvStr("SUI_DB_PATH", DefaultDBPath),
		LogLevel:    getEnvStr("SUI_LOG_LEVEL", DefaultLogLevel),
		LogFile:     getEnvStr("SUI_LOG_FILE", DefaultLogFile),
		LogMaxSize:  getEnvInt("SUI_LOG_MAX_SIZE", DefaultLogMaxSize),
		SecretKey:   getEnvStr("SUI_SECRET_KEY", ""),
	}
}

// getEnvStr retrieves a string environment variable or returns a default value.
func getEnvStr(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}

// getEnvInt retrieves an integer environment variable or returns a default value.
func getEnvInt(key string, defaultVal int) int {
	if val, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultVal
}

// GetVersion returns the current application version.
func GetVersion() string {
	return "0.0.1"
}

// IsDebug returns true if the application is running in debug mode.
func IsDebug() bool {
	return getEnvStr("SUI_DEBUG", "false") == "true"
}
