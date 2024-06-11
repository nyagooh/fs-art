package functions

import "testing"

func TestNonPrintable(t *testing.T) {
	
	tests := []struct {
		name string
		args string
		want string
	}{
		{"nonprintable","hello\\tworld","hello    world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonPrintable(tt.args); got != tt.want {
				t.Errorf("NonPrintable() = %v, want %v", got, tt.want)
			}
		})
	}
}
