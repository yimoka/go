// Package logger utils_test.go
package logger

import (
	"testing"
)

func TestToString(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"nil", nil, ""},
		{"string", "hello", "hello"},
		{"int", 42, "42"},
		{"int64", int64(123456789), "123456789"},
		{"uint64", uint64(987654321), "987654321"},
		{"float64", 3.14159, "3.14159"},
		{"float32", float32(2.718), "2.718"},
		{"bool_true", true, "true"},
		{"bool_false", false, "false"},
		{"bytes", []byte("test"), "test"},
		{"slice", []int{1, 2, 3}, "[1,2,3]"},
		{"map", map[string]string{"a": "b", "c": "d"}, `{"a":"b","c":"d"}`},
		{"struct", struct{ Name string }{"test"}, `{"Name":"test"}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := toString(tc.input)
			if result != tc.expected {
				t.Errorf("toString(%v) = %s, want %s", tc.input, result, tc.expected)
			}
		})
	}
}

// testStringer 实现了fmt.Stringer接口的类型
type testStringer struct {
	value string
}

func (ts testStringer) String() string {
	return "Stringer:" + ts.value
}

func TestToString_Stringer(t *testing.T) {
	// 测试实现了fmt.Stringer接口的类型
	result := toString(testStringer{value: "test"})
	expected := "Stringer:test"
	if result != expected {
		t.Errorf("toString(Stringer) = %s, want %s", result, expected)
	}
}

func TestToString_EdgeCases(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"empty_string", "", ""},
		{"zero_int", 0, "0"},
		{"zero_float", 0.0, "0"},
		{"negative_int", -42, "-42"},
		{"negative_float", -3.14, "-3.14"},
		{"large_int", 9223372036854775807, "9223372036854775807"},
		{"empty_slice", []int{}, "[]"},
		{"empty_map", map[string]int{}, "{}"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := toString(tc.input)
			if result != tc.expected {
				t.Errorf("toString(%v) = %s, want %s", tc.input, result, tc.expected)
			}
		})
	}
}
