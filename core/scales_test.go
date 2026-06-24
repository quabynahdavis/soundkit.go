package core

import (
	"math"
	"testing"
)

func TestGetScaleNotesMajor(t *testing.T) {
	cMajor, err := GetScaleNotes("C", "major", 4, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := []int{60, 62, 64, 65, 67, 69, 71}
	if !equalSlices(cMajor, expected) {
		t.Errorf("GetScaleNotes(C, major) = %v, want %v", cMajor, expected)
	}

	gMajor, err := GetScaleNotes("G", "major", 4, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedG := []int{67, 69, 71, 72, 74, 76, 77}
	if !equalSlices(gMajor, expectedG) {
		t.Errorf("GetScaleNotes(G, major) = %v, want %v", gMajor, expectedG)
	}
}

func TestGetScaleNotesMinor(t *testing.T) {
	aMinor, err := GetScaleNotes("A", "minor", 4, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := []int{69, 71, 72, 74, 76, 77, 79}
	if !equalSlices(aMinor, expected) {
		t.Errorf("GetScaleNotes(A, minor) = %v, want %v", aMinor, expected)
	}

	cMinor, err := GetScaleNotes("C", "natural_minor", 4, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedC := []int{60, 62, 63, 65, 67, 68, 70}
	if !equalSlices(cMinor, expectedC) {
		t.Errorf("GetScaleNotes(C, natural_minor) = %v, want %v", cMinor, expectedC)
	}
}

func TestGetScaleNotesMultipleOctaves(t *testing.T) {
	cMajor2Oct, err := GetScaleNotes("C", "major", 4, 2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(cMajor2Oct) != 14 {
		t.Errorf("Expected 14 notes for 2 octaves, got %d", len(cMajor2Oct))
	}
	if cMajor2Oct[0] != 60 {
		t.Errorf("cMajor2Oct[0] = %d, want 60", cMajor2Oct[0])
	}
	if cMajor2Oct[7] != 72 {
		t.Errorf("cMajor2Oct[7] = %d, want 72", cMajor2Oct[7])
	}
	if cMajor2Oct[13] != 83 {
		t.Errorf("cMajor2Oct[13] = %d, want 83", cMajor2Oct[13])
	}
}

func TestGetScaleNotesPentatonic(t *testing.T) {
	cPentatonicMajor, err := GetScaleNotes("C", "pentatonic_major", 4, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := []int{60, 62, 64, 67, 69}
	if !equalSlices(cPentatonicMajor, expected) {
		t.Errorf("GetScaleNotes(C, pentatonic_major) = %v, want %v", cPentatonicMajor, expected)
	}
}

func TestGetScaleNotesBlues(t *testing.T) {
	cBlues, err := GetScaleNotes("C", "blues", 4, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := []int{60, 63, 65, 66, 67, 70}
	if !equalSlices(cBlues, expected) {
		t.Errorf("GetScaleNotes(C, blues) = %v, want %v", cBlues, expected)
	}
}

func TestGetScaleNotesInvalid(t *testing.T) {
	if _, err := GetScaleNotes("C", "invalid_scale", 4, 1); err == nil {
		t.Error("Expected error for invalid scale type")
	}
	if _, err := GetScaleNotes("H", "major", 4, 1); err == nil {
		t.Error("Expected error for invalid scale root")
	}
}

func TestGetScaleFrequencies(t *testing.T) {
	cMajorFreq, err := GetScaleFrequencies("C", "major", 4, 1, 2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(cMajorFreq) != 7 {
		t.Fatalf("Expected 7 frequencies, got %d", len(cMajorFreq))
	}
	if math.Abs(cMajorFreq[0]-261.63) > 1e-2 {
		t.Errorf("cMajorFreq[0] = %f, want 261.63", cMajorFreq[0])
	}
	if math.Abs(cMajorFreq[6]-493.88) > 1e-2 {
		t.Errorf("cMajorFreq[6] = %f, want 493.88", cMajorFreq[6])
	}
}

func TestGetScaleNames(t *testing.T) {
	names := GetScaleNames()
	contains := func(list []string, target string) bool {
		for _, item := range list {
			if item == target {
				return true
			}
		}
		return false
	}
	if !contains(names, "major") {
		t.Error("Scale names should contain 'major'")
	}
	if !contains(names, "minor") {
		t.Error("Scale names should contain 'minor'")
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
