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

func TestGetEnvString(t *testing.T) {
	var tests = []struct {
		key           string
		expectedValue string
		expectedFound bool
	}{
		{StringName, StringValue, true},
		{StringName + "_nope", "", false},
	}

	for _, test := range tests {
		s, f := GetEnvString(test.key)
		if s != test.expectedValue {
			t.Errorf("Expected %s but got %s", test.expectedValue, s)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
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

func TestGetEnvInt(t *testing.T) {
	var tests = []struct {
		key           string
		expectedValue int
		expectedFound bool
	}{
		{IntName, IntValue, true},
		{IntName + "_nope", 0, false},
		{StringName, 0, false},
	}

	for _, test := range tests {
		i, f := GetEnvInt(test.key)
		if i != test.expectedValue {
			t.Errorf("Expected %d but got %d", test.expectedValue, i)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
		}
	}
}

func TestGetEnvBoolOrDefault(t *testing.T) {
	var tests = []struct {
		key           string
		defaultValue  bool
		expectedValue bool
	}{
		{BoolName, false, BoolValue},
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

func TestGetEnvBool(t *testing.T) {
	var tests = []struct {
		key           string
		expectedValue bool
		expectedFound bool
	}{
		{BoolName, BoolValue, true},
		{BoolName + "_nope", false, false},
		{StringName, false, false},
	}

	for _, test := range tests {
		v, f := GetEnvBool(test.key)
		if v != test.expectedValue {
			t.Errorf("Expected %v but got %v", test.expectedValue, v)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
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
func TestGetEnvDuration(t *testing.T) {
	var tests = []struct {
		key           string
		expectedValue time.Duration
		expectedFound bool
	}{
		{DurationName, DurationValue, true},
		{DurationName + "_nope", 0, false},
		{StringName, 0, false},
	}

	for _, test := range tests {
		d, f := GetEnvDuration(test.key)
		if d != test.expectedValue {
			t.Errorf("Expected %v but got %v", test.expectedValue, d)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
		}
	}
}
