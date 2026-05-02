package pkg

import (
	"fmt"
	"os"
)

func GetEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func GetEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		var val int
		fmt.Sscanf(v, "%d", &val)
		return val
	}
	return def
}