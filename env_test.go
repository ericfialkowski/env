package env

import (
	"fmt"
	"os"
	"testing"
	"time"
)

const (
	StringName    = "TESTING_STRING_ENV_VAR_1234"
	StringValue   = "Bob"
	BoolName      = "TESTING_BOOL_ENV_VAR_1234"
	BoolValue     = true
	IntName       = "TESTING_INT_ENV_VAR_1234"
	IntValue      = 1337
	DurationName  = "TESTING_DURATION_ENV_VAR_1234"
	DurationValue = time.Duration(1000 * 1000 * 1000 * 60)
)

func init() {
	os.Setenv(StringName, StringValue)
	os.Setenv(IntName, fmt.Sprint(IntValue))
	os.Setenv(BoolName, fmt.Sprint(BoolValue))
	os.Setenv(DurationName, fmt.Sprint(DurationValue))
}

func TestGetEnvStringOrDefault(t *testing.T) {
	var tests = []struct {
		key           string
		defaultValue  string
		expectedValue string
	}{
		{StringName, StringValue + "_nope", StringValue},
		{StringName + "_nope", StringValue, StringValue},
	}

	for _, test := range tests {
		s := GetEnvStringOrDefault(test.key, test.defaultValue)
		if s != test.expectedValue {
			t.Errorf("Expected %s but got %s", test.expectedValue, s)
		}
	}
}

func TestGetEnvIntOrDefault(t *testing.T) {
	var tests = []struct {
		key           string
		defaultValue  int
		expectedValue int
	}{
		{IntName, IntValue + 1, IntValue},
		{IntName + "_nope", IntValue, IntValue},
		{StringName, IntValue, IntValue},
	}

	for _, test := range tests {
		i := GetEnvIntOrDefault(test.key, test.defaultValue)
		if i != test.expectedValue {
			t.Errorf("Expected %d but got %d", test.expectedValue, i)
		}
	}
}

func TestGetEnvBoolOrDefault(t *testing.T) {
	var tests = []struct {
		key           string
		defaultValue  bool
		expectedValue bool
	}{
		{BoolName, !BoolValue, BoolValue},
		{BoolName + "_nope", BoolValue, BoolValue},
		{StringName, BoolValue, BoolValue},
	}

	for _, test := range tests {
		b := GetEnvBoolOrDefault(test.key, test.defaultValue)
		if b != test.expectedValue {
			t.Errorf("Expected %v but got %v", test.expectedValue, b)
		}
	}
}

func TestGetEnvDurationOrDefault(t *testing.T) {
	var tests = []struct {
		key           string
		defaultValue  time.Duration
		expectedValue time.Duration
	}{
		{DurationName, DurationValue - 1, DurationValue},
		{DurationName + "_nope", DurationValue, DurationValue},
		{StringName, DurationValue, DurationValue},
	}

	for _, test := range tests {
		d := GetEnvDurationOrDefault(test.key, test.defaultValue)
		if d != test.expectedValue {
			t.Errorf("Expected %v but got %v", test.expectedValue, d)
		}
	}
}
