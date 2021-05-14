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
// the supplied expectedValue if there is no environment variable named key.
func GetEnvStringOrDefault(key, defaultValue string) string {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	return envVal
}

// GetEnvString returns the same as os.LookupEnv. This is just a simple
// wrapper for completeness with the other parsing methods.
func GetEnvString(key string) (string, bool) {
	return os.LookupEnv(key)
}

// GetEnvBoolOrDefault returns the value in the system environment denoted by key as
// a bool or the supplied expectedValue if there is no environment variable named key or
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

// GetEnvBool returns the value in the system environment denoted by key as
// a bool and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a bool then (false, false) is returned
func GetEnvBool(key string) (bool, bool) {
	envVal, exists := os.LookupEnv(key)
	if !exists {
		return false, false
	}
	r, err := strconv.ParseBool(envVal)
	if err != nil {
		return false, false
	}
	return r, true
}

// GetEnvIntOrDefault returns the value in the system environment denoted by key as
// an int or the supplied expectedValue if there is no environment variable named key or
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

// GetEnvInt returns the value in the system environment denoted by key as
// an int and true. If there is no environment variable named key or
// if the value retrieved is not parsable as an int then (0, false) is returned
func GetEnvInt(key string) (int, bool) {
	envVal, exists := os.LookupEnv(key)
	if !exists {
		return 0, false
	}
	r, err := strconv.Atoi(envVal)
	if err != nil {
		return 0, false
	}
	return r, true
}

// GetEnvDurationOrDefault returns the value in the system environment denoted by key as
// a time.Duration or the supplied expectedValue if there is no environment variable named key or
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

// GetEnvDuration returns the value in the system environment denoted by key as
// a time.Duration and true. If there is no environment variable named key or
// if the value retrieved is not parsable as an time.Duration then (0, false) is returned
// See time.ParseDuration() for the allowed formatting of the environment variable.
func GetEnvDuration(key string) (time.Duration, bool) {
	envVal, exists := os.LookupEnv(key)
	if !exists {
		return 0, false
	}
	r, err := time.ParseDuration(envVal)
	if err != nil {
		return 0, false
	}
	return r, true
}

// GetEnvFloat32OrDefault returns the value in the system environment denoted by key as
// an int or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as a float32.
func GetEnvFloat32OrDefault(key string, defaultValue float32) float32 {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return defaultValue
	}
	return float32(r)
}

// GetEnvFloat32 returns the value in the system environment denoted by key as
// an int and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a float32 then (0, false) is returned
func GetEnvFloat32(key string) (float32, bool) {
	envVal, exists := os.LookupEnv(key)
	if !exists {
		return 0, false
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return 0, false
	}
	return float32(r), true
}

// GetEnvFloat64OrDefault returns the value in the system environment denoted by key as
// an int or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as a float64.
func GetEnvFloat64OrDefault(key string, defaultValue float64) float64 {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return defaultValue
	}
	return r
}

// GetEnvFloat64 returns the value in the system environment denoted by key as
// an int and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a float64 then (0, false) is returned
func GetEnvFloat64(key string) (float64, bool) {
	envVal, exists := os.LookupEnv(key)
	if !exists {
		return 0, false
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return 0, false
	}
	return r, true
}
