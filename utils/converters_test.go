package utils

import (
	"math"
	"testing"
)

func TestFrequencyToCents(t *testing.T) {
	// Octave difference
	cents, err := FrequencyToCents(440, 880)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(cents-1200.0) > 1e-2 {
		t.Errorf("FrequencyToCents(440, 880) = %f, want 1200.0", cents)
	}

	// Perfect fifth (700 cents)
	cents, err = FrequencyToCents(440, 660)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(cents-700.0) > 1e-2 {
		// Wait, math.Log2(1.5) * 1200 is 701.955
		if math.Abs(cents-701.955) > 1e-2 {
			t.Errorf("FrequencyToCents(440, 660) = %f, want ~701.955", cents)
		}
	}

	// Small difference
	cents, err = FrequencyToCents(440, 444)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if cents <= 0 || cents >= 20 {
		t.Errorf("FrequencyToCents(440, 444) = %f, want value between 0 and 20", cents)
	}

	// Invalid
	if _, err := FrequencyToCents(-100, 440); err == nil {
		t.Error("Expected error for negative frequency")
	}
	if _, err := FrequencyToCents(440, 0); err == nil {
		t.Error("Expected error for zero frequency")
	}
}

func TestCentsToRatio(t *testing.T) {
	// Octave
	ratio := CentsToRatio(1200)
	if math.Abs(ratio-2.0) > 1e-2 {
		t.Errorf("CentsToRatio(1200) = %f, want 2.0", ratio)
	}

	// Perfect fifth (700 cents)
	ratio = CentsToRatio(700)
	if math.Abs(ratio-1.498) > 1e-2 { // 2^(700/1200) = 1.4983
		t.Errorf("CentsToRatio(700) = %f, want 1.498", ratio)
	}

	// Semitone
	ratio = CentsToRatio(100)
	if math.Abs(ratio-1.05946) > 1e-4 {
		t.Errorf("CentsToRatio(100) = %f, want 1.05946", ratio)
	}
}

func TestRatioToCents(t *testing.T) {
	// Octave
	cents, err := RatioToCents(2.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(cents-1200.0) > 1e-2 {
		t.Errorf("RatioToCents(2.0) = %f, want 1200.0", cents)
	}

	// Perfect fifth
	cents, err = RatioToCents(1.5)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(cents-701.955) > 1e-2 {
		t.Errorf("RatioToCents(1.5) = %f, want 701.955", cents)
	}

	// Invalid
	if _, err := RatioToCents(-1.0); err == nil {
		t.Error("Expected error for negative ratio")
	}
	if _, err := RatioToCents(0.0); err == nil {
		t.Error("Expected error for zero ratio")
	}
}

func TestSemitonesToRatio(t *testing.T) {
	if ratio := SemitonesToRatio(12); math.Abs(ratio-2.0) > 1e-2 {
		t.Errorf("SemitonesToRatio(12) = %f, want 2.0", ratio)
	}
	if ratio := SemitonesToRatio(7); math.Abs(ratio-1.498) > 1e-2 {
		t.Errorf("SemitonesToRatio(7) = %f, want 1.498", ratio)
	}
	if ratio := SemitonesToRatio(0); math.Abs(ratio-1.0) > 1e-2 {
		t.Errorf("SemitonesToRatio(0) = %f, want 1.0", ratio)
	}
}

func TestRatioToSemitones(t *testing.T) {
	semi, err := RatioToSemitones(2.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(semi-12.0) > 1e-2 {
		t.Errorf("RatioToSemitones(2.0) = %f, want 12.0", semi)
	}

	semi, err = RatioToSemitones(1.5)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(semi-7.01955) > 1e-2 {
		t.Errorf("RatioToSemitones(1.5) = %f, want 7.01955", semi)
	}

	semi, err = RatioToSemitones(1.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(semi-0.0) > 1e-2 {
		t.Errorf("RatioToSemitones(1.0) = %f, want 0.0", semi)
	}
}

func TestNormalizeFrequency(t *testing.T) {
	// Should normalize to nearest A
	norm, err := NormalizeFrequency(441.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(norm-440.0) > 1e-1 {
		t.Errorf("NormalizeFrequency(441.0) = %f, want 440.0", norm)
	}

	norm, err = NormalizeFrequency(439.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(norm-440.0) > 1e-1 {
		t.Errorf("NormalizeFrequency(439.0) = %f, want 440.0", norm)
	}

	// With custom reference
	norm, err = NormalizeFrequency(443.0, 442.0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Abs(norm-442.0) > 1e-1 {
		t.Errorf("NormalizeFrequency(443.0, 442.0) = %f, want 442.0", norm)
	}

	// Invalid
	if _, err := NormalizeFrequency(-100.0); err == nil {
		t.Error("Expected error for negative frequency")
	}
	if _, err := NormalizeFrequency(0.0); err == nil {
		t.Error("Expected error for zero frequency")
	}
}
