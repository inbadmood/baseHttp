package utils

import (
	"testing"
)

func TestInterfaceToString(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{"int", 123, "123"},
		{"string", "hello", "hello"},
		{"float64", 3.14, "3.14"},
		{"bool", true, "true"},
		{"struct", struct{}{}, "{}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceToString(tt.input); got != tt.want {
				t.Errorf("InterfaceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
