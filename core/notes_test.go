package core

import (
	"math"
	"testing"
)

func TestMidiKey(t *testing.T) {
	// Basic note to MIDI conversion
	tests := []struct {
		input string
		want  int
	}{
		{"C4", 60},
		{"A4", 69},
		{"C0", 12},
		{"C-1", 0},
		{"G9", 127},
		// Sharps
		{"C#4", 61},
		{"D#4", 63},
		{"F#4", 66},
		{"G#4", 68},
		{"A#4", 70},
		// Flats
		{"Db4", 61},
		{"Eb4", 63},
		{"Gb4", 66},
		{"Ab4", 68},
		{"Bb4", 70},
		// Case insensitive
		{"c4", 60},
		{"db4", 61},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := MidiKey(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error for %s: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("MidiKey(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestMidiKeyInvalid(t *testing.T) {
	invalidNotes := []string{"H4", "C", "C#4extra"}
	for _, n := range invalidNotes {
		if _, err := MidiKey(n); err == nil || !errorsIs(err, ErrInvalidNote) {
			t.Errorf("Expected ErrInvalidNote for input %q, got: %v", n, err)
		}
	}

	invalidOctaves := []string{"C11", "C-2"}
	for _, n := range invalidOctaves {
		if _, err := MidiKey(n); err == nil || !errorsIs(err, ErrInvalidOctave) {
			t.Errorf("Expected ErrInvalidOctave for input %q, got: %v", n, err)
		}
	}
}

func errorsIs(err error, target error) bool {
	if err == nil {
		return target == nil
	}
	// Direct comparison or unwrap
	if err == target {
		return true
	}
	type unwrapper interface {
		Unwrap() error
	}
	if u, ok := err.(unwrapper); ok {
		return errorsIs(u.Unwrap(), target)
	}
	return false
}

func TestMidiFreq(t *testing.T) {
	// A4
	freq, err := MidiFreq("A4", 2, 440.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(freq-440.0) > 1e-2 {
		t.Errorf("MidiFreq(A4) = %f, want 440.0", freq)
	}

	// C4
	freq, err = MidiFreq("C4", 2, 440.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(freq-261.63) > 1e-2 {
		t.Errorf("MidiFreq(C4) = %f, want 261.63", freq)
	}

	// A3
	freq, err = MidiFreq("A3", 2, 440.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(freq-220.0) > 1e-2 {
		t.Errorf("MidiFreq(A3) = %f, want 220.0", freq)
	}

	// Rounding
	freq, err = MidiFreq("C4", -1, 440.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// Check it returns a raw float with full precision
	if math.Abs(freq-261.625565) > 1e-5 {
		t.Errorf("MidiFreq(C4, -1) should not round, got %f", freq)
	}

	freq, err = MidiFreq("C4", 4, 440.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(freq-261.6256) > 1e-5 {
		t.Errorf("MidiFreq(C4, 4) = %f, want 261.6256", freq)
	}
}

func TestFreqToMidi(t *testing.T) {
	tests := []struct {
		freq float64
		want int
	}{
		{440.0, 69},
		{261.63, 60},
		{880.0, 81},
	}

	for _, tt := range tests {
		got, err := FreqToMidi(tt.freq, 440.0)
		if err != nil {
			t.Fatalf("Unexpected error for %f: %v", tt.freq, err)
		}
		if got != tt.want {
			t.Errorf("FreqToMidi(%f) = %d, want %d", tt.freq, got, tt.want)
		}
	}
}

func TestMidiToNoteName(t *testing.T) {
	tests := []struct {
		midi      int
		useSharps bool
		want      string
	}{
		{60, true, "C4"},
		{69, true, "A4"},
		{0, true, "C-1"},
		{127, true, "G9"},
		{61, false, "Db4"},
		{63, false, "Eb4"},
		{66, false, "Gb4"},
	}

	for _, tt := range tests {
		got, err := MidiToNoteName(tt.midi, tt.useSharps)
		if err != nil {
			t.Fatalf("Unexpected error for %d: %v", tt.midi, err)
		}
		if got != tt.want {
			t.Errorf("MidiToNoteName(%d, %t) = %s, want %s", tt.midi, tt.useSharps, got, tt.want)
		}
	}
}

func TestNotesToMidi(t *testing.T) {
	notes := []string{"C4", "E4", "G4", "A4", "invalid_note"}
	result := NotesToMidi(notes)

	if len(result) != 5 {
		t.Fatalf("Expected 5 results, got %d", len(result))
	}
	if val, ok := result[0].(int); !ok || val != 60 {
		t.Errorf("result[0] = %v, want 60", result[0])
	}
	if val, ok := result[1].(int); !ok || val != 64 {
		t.Errorf("result[1] = %v, want 64", result[1])
	}
	if val, ok := result[2].(int); !ok || val != 67 {
		t.Errorf("result[2] = %v, want 67", result[2])
	}
	if val, ok := result[3].(int); !ok || val != 69 {
		t.Errorf("result[3] = %v, want 69", result[3])
	}
	if _, ok := result[4].(string); !ok {
		t.Errorf("result[4] = %v, want string error message", result[4])
	}
}
