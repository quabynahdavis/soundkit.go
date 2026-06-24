package utils

import (
	"testing"
)

func TestValidateNoteName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"C4", "C4", true},
		{"C#4", "C#4", true},
		{"Db4", "Db4", true},
		{"A-1", "A-1", true},
		{"G9", "G9", true},
		{"H4", "H4", false},
		{"C", "C", false},
		{"4", "4", false},
		{"C4extra", "C4extra", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateNoteName(tt.input); got != tt.want {
				t.Errorf("ValidateNoteName(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestValidateMidiRange(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"0", 0, true},
		{"60", 60, true},
		{"127", 127, true},
		{"-1", -1, false},
		{"128", 128, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateMidiRange(tt.input); got != tt.want {
				t.Errorf("ValidateMidiRange(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestValidateFrequency(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  bool
	}{
		{"440", 440.0, true},
		{"20", 20.0, true},
		{"20000", 20000.0, true},
		{"-100", -100.0, false},
		{"0", 0.0, false},
		{"30000", 30000.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateFrequency(tt.input); got != tt.want {
				t.Errorf("ValidateFrequency(%f) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestValidateOctave(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"-1", -1, true},
		{"0", 0, true},
		{"4", 4, true},
		{"10", 10, true},
		{"-2", -2, false},
		{"11", 11, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateOctave(tt.input); got != tt.want {
				t.Errorf("ValidateOctave(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestNormalizeNoteName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"c4", "c4", "C4"},
		{"C#4", "C#4", "C#4"},
		{"db4", "db4", "DB4"},
		{"C ♯4", "C ♯4", "C#4"},
		{"C ♭4", "C ♭4", "CB4"},
		{"C-4", "C-4", "C-4"},
		{"C 4", "C 4", "C4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeNoteName(tt.input); got != tt.want {
				t.Errorf("NormalizeNoteName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
