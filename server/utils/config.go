package utils

import "os"

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var (
	logFilePath = getEnv("LOG_FILE_PATH", "logs/zxqpro.log")
	useLogFile  = getEnv("USE_LOG_FILE", "0")
	CookieName  = "zxqpro_cookie"
)
