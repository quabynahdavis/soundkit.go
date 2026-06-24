package soundkit_test

import (
	"math"
	"testing"

	"github.com/quabynahdavis/soundkit.go"
)

func TestFacadePublicAPI(t *testing.T) {
	// 1. Constants and Errors
	if soundkit.ConcertPitch != 440 {
		t.Errorf("Expected ConcertPitch to be 440, got %d", soundkit.ConcertPitch)
	}
	if soundkit.ErrInvalidNote == nil {
		t.Error("Expected ErrInvalidNote to be defined")
	}

	// 2. Note conversion and validation
	midi, err := soundkit.MidiKey("C4")
	if err != nil {
		t.Fatalf("Unexpected error for MidiKey: %v", err)
	}
	if midi != 60 {
		t.Errorf("MidiKey(\"C4\") = %d, want 60", midi)
	}

	freq, err := soundkit.MidiFreq("A4", 2, float64(soundkit.ConcertPitch))
	if err != nil {
		t.Fatalf("Unexpected error for MidiFreq: %v", err)
	}
	if freq != 440.0 {
		t.Errorf("MidiFreq(\"A4\") = %f, want 440.0", freq)
	}

	// 3. Validators & Converters
	if !soundkit.ValidateNoteName("C#4") {
		t.Error("Expected C#4 to be a valid note name")
	}

	cents, err := soundkit.FrequencyToCents(440.0, 880.0)
	if err != nil {
		t.Fatalf("Unexpected error for FrequencyToCents: %v", err)
	}
	if math.Abs(cents-1200.0) > 1e-2 {
		t.Errorf("FrequencyToCents(440, 880) = %f, want 1200.0", cents)
	}

	// 4. Chords and Scales
	chordNotes, err := soundkit.GetChordNotes("C", "maj", 4, 0)
	if err != nil {
		t.Fatalf("Unexpected error for GetChordNotes: %v", err)
	}
	if len(chordNotes) != 3 || chordNotes[0] != 60 || chordNotes[1] != 64 || chordNotes[2] != 67 {
		t.Errorf("GetChordNotes(C, maj) = %v, want [60 64 67]", chordNotes)
	}

	scaleNotes, err := soundkit.GetScaleNotes("C", "major", 4, 1)
	if err != nil {
		t.Fatalf("Unexpected error for GetScaleNotes: %v", err)
	}
	if len(scaleNotes) != 7 || scaleNotes[0] != 60 || scaleNotes[6] != 71 {
		t.Errorf("GetScaleNotes(C, major) = %v, want [60 62 64 65 67 69 71]", scaleNotes)
	}
}
