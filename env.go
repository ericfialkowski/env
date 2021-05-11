// Package env provides simple environment variable retrieval and parsing into builtin
// types including providing the default value if the environment variable is not set
// or not in the correct format.
package env

import (
	"os"
	"strconv"
	"time"
)

// GetEnvStringOrDefault returns the value in the system environment denoted by key or
// the supplied defaultValue if there is no environment variable named key.
func GetEnvStringOrDefault(key, defaultValue string) string {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	return envVal
}

// GetEnvBoolOrDefault returns the value in the system environment denoted by key as
// a bool or the supplied defaultValue if there is no environment variable named key or
// if the value retrieved is not parsable as a bool.
func GetEnvBoolOrDefault(key string, defaultValue bool) bool {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	r, err := strconv.ParseBool(envVal)
	if err != nil {
		return defaultValue
	}
	return r
}

// GetEnvIntOrDefault returns the value in the system environment denoted by key as
// an int or the supplied defaultValue if there is no environment variable named key or
// if the value retrieved is not parsable as an int.
func GetEnvIntOrDefault(key string, defaultValue int) int {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	r, err := strconv.Atoi(envVal)
	if err != nil {
		return defaultValue
	}
	return r
}

// GetEnvDurationOrDefault returns the value in the system environment denoted by key as
// a time.Duration or the supplied defaultValue if there is no environment variable named key or
// if the value retrieved is not parsable as a time.Duration. See time.ParseDuration() for the
// allowed formatting of the environment variable.
func GetEnvDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	r, err := time.ParseDuration(envVal)
	if err != nil {
		return defaultValue
	}
	return r
}
