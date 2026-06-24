package soundkit

import (
	"errors"
	"fmt"
	"math"
)

// FrequencyToCents converts a frequency ratio to cents.
func FrequencyToCents(freq1, freq2 float64) (float64, error) {
	if freq1 <= 0 || freq2 <= 0 {
		return 0, fmt.Errorf("%w: frequencies must be positive", ErrInvalidFrequency)
	}
	return 1200 * math.Log2(freq2/freq1), nil
}

// CentsToRatio converts cents to a frequency ratio.
func CentsToRatio(cents float64) float64 {
	return math.Pow(2, cents/1200)
}

// RatioToCents converts a frequency ratio to cents.
func RatioToCents(ratio float64) (float64, error) {
	if ratio <= 0 {
		return 0, fmt.Errorf("%w: ratio must be positive", ErrInvalidFrequency)
	}
	return 1200 * math.Log2(ratio), nil
}

// SemitonesToRatio converts semitones to a frequency ratio.
func SemitonesToRatio(semitones float64) float64 {
	return math.Pow(2, semitones/12)
}

// RatioToSemitones converts a frequency ratio to semitones.
func RatioToSemitones(ratio float64) (float64, error) {
	if ratio <= 0 {
		return 0, fmt.Errorf("%w: ratio must be positive", ErrInvalidFrequency)
	}
	return 12 * math.Log2(ratio), nil
}

// NormalizeFrequency normalizes a frequency to the nearest reference pitch.
// Note: Go doesn't support default arguments, so you must explicitly pass the reference pitch.
func NormalizeFrequency(frequency, reference float64) (float64, error) {
	if frequency <= 0 {
		return 0, fmt.Errorf("%w: frequency must be positive", ErrInvalidFrequency)
	}

	// In Go, math.Round returns a float64 directly
	octaveOffset := math.Round(12 * math.Log2(frequency/reference)) / 12
	nearestA := reference * math.Pow(2, octaveOffset)
	
	return nearestA, nil
}