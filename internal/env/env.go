package env

import (
	"fmt"
	"os"
	"strings"
)

const (
	EnvMongoURL = "MONGO_URL"
)

func GetMongoURL() string {
	return GetEnv(EnvMongoURL)
}

func GetEnv(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return value
}

func CheckEnvExist(keys ...string) error {
	var err []string

	for _, key := range keys {
		if os.Getenv(key) == "" {
			err = append(err, key)
		}
	}

	if len(err) > 0 {
		return fmt.Errorf("variables are required: %v", err)
	}

	return nil
}
