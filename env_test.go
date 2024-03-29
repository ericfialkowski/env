package env

import (
	"fmt"
	"os"
	"reflect"
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
	Float32Name   = "TESTING_F32_ENV_VAR_1234"
	Float32Value  = float32(1337.7331)
	Float64Name   = "TESTING_F64_ENV_VAR_1234"
	Float64Value  = float64(1337.7331)
)

func init() {
	_ = os.Setenv(StringName, StringValue)
	_ = os.Setenv(IntName, fmt.Sprint(IntValue))
	_ = os.Setenv(BoolName, fmt.Sprint(BoolValue))
	_ = os.Setenv(DurationName, fmt.Sprint(DurationValue))
	_ = os.Setenv(Float32Name, fmt.Sprint(Float32Value))
	_ = os.Setenv(Float64Name, fmt.Sprint(Float64Value))
}

func TestGetStringOrDefault(t *testing.T) {
	var tests = []struct {
		key           string
		defaultValue  string
		expectedValue string
	}{
		{StringName, StringValue + "_nope", StringValue},
		{StringName + "_nope", StringValue, StringValue},
	}

	for _, test := range tests {
		s := StringOrDefault(test.key, test.defaultValue)
		if s != test.expectedValue {
			t.Errorf("Expected %s but got %s", test.expectedValue, s)
		}
	}
}

func TestGetStringOrPanic(t *testing.T) {
	var tests = []struct {
		key           string
		expectedValue string
		panicked      bool
	}{
		{StringName, StringValue, false},
		{StringName + "_nope", "", true},
	}

	var panicOccurred bool
	defer func() {
		if err := recover(); err != nil {
			panicOccurred = true
		}
	}()
	for _, test := range tests {
		panicOccurred = false
		s := MustString(test.key)
		if s != test.expectedValue {
			t.Errorf("Expected %s but got %s", test.expectedValue, s)
		}
		if panicOccurred != test.panicked {
			if panicOccurred {
				t.Error("Panicked where we shouldn't have")
			} else {
				t.Error("Didn't panic but should have")
			}
		}
	}
}

func TestGetString(t *testing.T) {
	var tests = []struct {
		key           string
		expectedValue string
		expectedFound bool
	}{
		{StringName, StringValue, true},
		{StringName + "_nope", "", false},
	}

	for _, test := range tests {
		s, f := String(test.key)
		if s != test.expectedValue {
			t.Errorf("Expected %s but got %s", test.expectedValue, s)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
		}
	}
}

func TestGetIntOrDefault(t *testing.T) {
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
		i := IntOrDefault(test.key, test.defaultValue)
		if i != test.expectedValue {
			t.Errorf("Expected %d but got %d", test.expectedValue, i)
		}
	}
}

func TestGetInt(t *testing.T) {
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
		i, f := Int(test.key)
		if i != test.expectedValue {
			t.Errorf("Expected %d but got %d", test.expectedValue, i)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
		}
	}
}

func TestGetBoolOrDefault(t *testing.T) {
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
		b := BoolOrDefault(test.key, test.defaultValue)
		if b != test.expectedValue {
			t.Errorf("Expected %v but got %v", test.expectedValue, b)
		}
	}
}

func TestGetBool(t *testing.T) {
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
		v, f := Bool(test.key)
		if v != test.expectedValue {
			t.Errorf("Expected %v but got %v", test.expectedValue, v)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
		}
	}
}

func TestGetDurationOrDefault(t *testing.T) {
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
		d := DurationOrDefault(test.key, test.defaultValue)
		if d != test.expectedValue {
			t.Errorf("Expected %v but got %v", test.expectedValue, d)
		}
	}
}
func TestGetDuration(t *testing.T) {
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
		d, f := Duration(test.key)
		if d != test.expectedValue {
			t.Errorf("Expected %v but got %v", test.expectedValue, d)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
		}
	}
}

func TestGetFloat32OrDefault(t *testing.T) {
	var tests = []struct {
		key           string
		defaultValue  float32
		expectedValue float32
	}{
		{Float32Name, Float32Value + 1, Float32Value},
		{Float32Name + "_nope", Float32Value, Float32Value},
		{StringName, Float32Value, Float32Value},
	}

	for _, test := range tests {
		i := Float32OrDefault(test.key, test.defaultValue)
		if i != test.expectedValue {
			t.Errorf("Expected %f but got %f", test.expectedValue, i)
		}
		if reflect.TypeOf(i).String() != "float32" {
			t.Errorf("Expected float32 but got %v", reflect.TypeOf(i))
		}
	}
}

func TestGetFloat32(t *testing.T) {
	var tests = []struct {
		key           string
		expectedValue float32
		expectedFound bool
	}{
		{Float32Name, Float32Value, true},
		{Float32Name + "_nope", 0, false},
		{StringName, 0, false},
	}

	for _, test := range tests {
		i, f := Float32(test.key)
		if i != test.expectedValue {
			t.Errorf("Expected %f but got %f", test.expectedValue, i)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
		}
	}
}

func TestGetFloat64OrDefault(t *testing.T) {
	var tests = []struct {
		key           string
		defaultValue  float64
		expectedValue float64
	}{
		{Float64Name, Float64Value + 1, Float64Value},
		{Float64Name + "_nope", Float64Value, Float64Value},
		{StringName, Float64Value, Float64Value},
	}

	for _, test := range tests {
		i := Float64OrDefault(test.key, test.defaultValue)
		if i != test.expectedValue {
			t.Errorf("Expected %f but got %f", test.expectedValue, i)
		}
	}
}

func TestGetFloat64(t *testing.T) {
	var tests = []struct {
		key           string
		expectedValue float64
		expectedFound bool
	}{
		{Float64Name, Float64Value, true},
		{Float64Name + "_nope", 0, false},
		{StringName, 0, false},
	}

	for _, test := range tests {
		i, f := Float64(test.key)
		if i != test.expectedValue {
			t.Errorf("Expected %f but got %f", test.expectedValue, i)
		}
		if f != test.expectedFound {
			t.Errorf("Expected found of %v but got %v", test.expectedFound, f)
		}
	}
}
