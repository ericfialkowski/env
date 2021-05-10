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

func TestGetEnvString(t *testing.T) {
	s := GetEnvStringOrDefault(StringName, StringValue+"_nope")
	if s != StringValue {
		t.Errorf("Expected %s but got %s", StringValue, s)
	}
}

func TestGetEnvStringDefault(t *testing.T) {
	s := GetEnvStringOrDefault(StringName+"_nope", StringValue)
	if s != StringValue {
		t.Errorf("Expected %s but got %s", StringValue, s)
	}
}

func TestGetEnvInt(t *testing.T) {
	i := GetEnvIntOrDefault(IntName, IntValue-1)
	if i != IntValue {
		t.Errorf("Expected %d but got %d", IntValue, i)
	}
}

func TestGetEnvIntDefault(t *testing.T) {
	i := GetEnvIntOrDefault(IntName+"_nope", IntValue)
	if i != IntValue {
		t.Errorf("Expected %d but got %d", IntValue, i)
	}
}

func TestGetEnvIntDefaultFromBadValue(t *testing.T) {
	i := GetEnvIntOrDefault(StringName, IntValue)
	if i != IntValue {
		t.Errorf("Expected %d but got %d", IntValue, i)
	}
}

func TestGetEnvBool(t *testing.T) {
	b := GetEnvBoolOrDefault(BoolName, !BoolValue)
	if b != BoolValue {
		t.Errorf("Expected %v but got %v", BoolValue, b)
	}
}

func TestGetEnvBoolOrDefaultGetsDefault(t *testing.T) {
	b := GetEnvBoolOrDefault(BoolName+"_nope", BoolValue)
	if b != BoolValue {
		t.Errorf("Expected %v but got %v", BoolValue, b)
	}
}

func TestGetEnvBoolOrDefaultGetsDefaultFromBadValue(t *testing.T) {
	b := GetEnvBoolOrDefault(StringName, BoolValue)
	if b != BoolValue {
		t.Errorf("Expected %v but got %v", BoolValue, b)
	}
}

func TestGetEnvDurationOrDefault(t *testing.T) {
	d := GetEnvDurationOrDefault(DurationName, DurationValue-1)
	if d != DurationValue {
		t.Errorf("Expected %v but got %v", DurationValue, d)
	}
}

func TestGetEnvDurationOrDefaultGetsDefault(t *testing.T) {
	d := GetEnvDurationOrDefault(DurationName+"_nope", DurationValue)
	if d != DurationValue {
		t.Errorf("Expected %v but got %v", DurationValue, d)
	}
}

func TestGetEnvDurationOrDefaultGetsDefaultFromBadValue(t *testing.T) {
	d := GetEnvDurationOrDefault(StringName, DurationValue)
	if d != DurationValue {
		t.Errorf("Expected %v but got %v", DurationValue, d)
	}
}
