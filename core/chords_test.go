package core

import (
	"math"
	"testing"
)

func TestGetChordNotesMajor(t *testing.T) {
	cMajor, err := GetChordNotes("C", "maj", 4, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := []int{60, 64, 67}
	if !equalSlices(cMajor, expected) {
		t.Errorf("GetChordNotes(C, maj) = %v, want %v", cMajor, expected)
	}

	gMajor, err := GetChordNotes("G", "major", 4, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedG := []int{67, 71, 74}
	if !equalSlices(gMajor, expectedG) {
		t.Errorf("GetChordNotes(G, major) = %v, want %v", gMajor, expectedG)
	}
}

func TestGetChordNotesMinor(t *testing.T) {
	cMinor, err := GetChordNotes("C", "min", 4, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := []int{60, 63, 67}
	if !equalSlices(cMinor, expected) {
		t.Errorf("GetChordNotes(C, min) = %v, want %v", cMinor, expected)
	}
}

func TestGetChordNotesSeventh(t *testing.T) {
	cMajor7, err := GetChordNotes("C", "maj7", 4, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := []int{60, 64, 67, 71}
	if !equalSlices(cMajor7, expected) {
		t.Errorf("GetChordNotes(C, maj7) = %v, want %v", cMajor7, expected)
	}
}

func TestGetChordNotesInversions(t *testing.T) {
	// First inversion
	cMajorFirst, err := GetChordNotes("C", "maj", 4, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedFirst := []int{64, 67, 72}
	if !equalSlices(cMajorFirst, expectedFirst) {
		t.Errorf("GetChordNotes(C, maj, inversion=1) = %v, want %v", cMajorFirst, expectedFirst)
	}

	// Second inversion
	cMajorSecond, err := GetChordNotes("C", "maj", 4, 2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedSecond := []int{67, 72, 76}
	if !equalSlices(cMajorSecond, expectedSecond) {
		t.Errorf("GetChordNotes(C, maj, inversion=2) = %v, want %v", cMajorSecond, expectedSecond)
	}
}

func TestGetChordFrequencies(t *testing.T) {
	cMajorFreq, err := GetChordFrequencies("C", "maj", 4, 2, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(cMajorFreq) != 3 {
		t.Fatalf("Expected 3 frequencies, got %d", len(cMajorFreq))
	}
	if math.Abs(cMajorFreq[0]-261.63) > 1e-2 {
		t.Errorf("cMajorFreq[0] = %f, want 261.63", cMajorFreq[0])
	}
	if math.Abs(cMajorFreq[1]-329.63) > 1e-2 {
		t.Errorf("cMajorFreq[1] = %f, want 329.63", cMajorFreq[1])
	}
	if math.Abs(cMajorFreq[2]-392.00) > 1e-2 {
		t.Errorf("cMajorFreq[2] = %f, want 392.00", cMajorFreq[2])
	}
}
