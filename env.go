// Package env provides simple environment variable retrieval and parsing into builtin
// types including providing the default value if the environment variable is not set
// or not in the correct format.
package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// GetStringOrDefault returns the value in the system environment denoted by key or
// the supplied expectedValue if there is no environment variable named key.
func GetStringOrDefault(key, defaultValue string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return envVal
}

// GetStringOrPanic returns the value in the system environment denoted by key or
// panics. Should really only be used when you can't run without the value and there
// is no viable default to provide
func GetStringOrPanic(key string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Missing required environment variable value for %s", key))
	}
	return envVal
}

// GetString returns the same as os.LookupEnv. This is just a simple
// wrapper for completeness with the other parsing methods.
func GetString(key string) (string, bool) {
	return os.LookupEnv(key)
}

// GetBoolOrDefault returns the value in the system environment denoted by key as
// a bool or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as a bool.
func GetBoolOrDefault(key string, defaultValue bool) bool {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	r, err := strconv.ParseBool(envVal)
	if err != nil {
		return defaultValue
	}
	return r
}

// GetBool returns the value in the system environment denoted by key as
// a bool and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a bool then (false, false) is returned
func GetBool(key string) (bool, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return false, false
	}
	r, err := strconv.ParseBool(envVal)
	if err != nil {
		return false, false
	}
	return r, true
}

// GetIntOrDefault returns the value in the system environment denoted by key as
// an int or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as an int.
func GetIntOrDefault(key string, defaultValue int) int {
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

// GetInt returns the value in the system environment denoted by key as
// an int and true. If there is no environment variable named key or
// if the value retrieved is not parsable as an int then (0, false) is returned
func GetInt(key string) (int, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return 0, false
	}
	r, err := strconv.Atoi(envVal)
	if err != nil {
		return 0, false
	}
	return r, true
}

// GetDurationOrDefault returns the value in the system environment denoted by key as
// a time.Duration or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as a time.Duration. See time.ParseDuration() for the
// allowed formatting of the environment variable.
func GetDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
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

// GetDuration returns the value in the system environment denoted by key as
// a time.Duration and true. If there is no environment variable named key or
// if the value retrieved is not parsable as an time.Duration then (0, false) is returned
// See time.ParseDuration() for the allowed formatting of the environment variable.
func GetDuration(key string) (time.Duration, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
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

// GetFloat32 returns the value in the system environment denoted by key as
// an int and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a float32 then (0, false) is returned
func GetFloat32(key string) (float32, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return 0, false
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return 0, false
	}
	return float32(r), true
}

// GetFloat64OrDefault returns the value in the system environment denoted by key as
// an int or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as a float64.
func GetFloat64OrDefault(key string, defaultValue float64) float64 {
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

// GetFloat64 returns the value in the system environment denoted by key as
// an int and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a float64 then (0, false) is returned
func GetFloat64(key string) (float64, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return 0, false
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return 0, false
	}
	return r, true
}
