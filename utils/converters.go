package utils

import (
	"errors"
	"fmt"
	"math"
)

// ConcertPitch is the standard tuning frequency (440Hz for Key A).
const ConcertPitch = 440

// ErrInvalidFrequency is raised when an invalid frequency is provided.
var ErrInvalidFrequency = errors.New("invalid frequency is provided")

// FrequencyToCents converts frequency ratio to cents.
func FrequencyToCents(freq1, freq2 float64) (float64, error) {
	if freq1 <= 0 || freq2 <= 0 {
		return 0, fmt.Errorf("%w: frequencies must be positive", ErrInvalidFrequency)
	}
	return 1200 * math.Log2(freq2/freq1), nil
}

// CentsToRatio converts cents to frequency ratio.
func CentsToRatio(cents float64) float64 {
	return math.Pow(2, cents/1200)
}

// RatioToCents converts frequency ratio to cents.
func RatioToCents(ratio float64) (float64, error) {
	if ratio <= 0 {
		return 0, fmt.Errorf("%w: ratio must be positive", ErrInvalidFrequency)
	}
	return 1200 * math.Log2(ratio), nil
}

// SemitonesToRatio converts semitones to frequency ratio.
func SemitonesToRatio(semitones float64) float64 {
	return math.Pow(2, semitones/12)
}

// RatioToSemitones converts frequency ratio to semitones.
func RatioToSemitones(ratio float64) (float64, error) {
	if ratio <= 0 {
		return 0, fmt.Errorf("%w: ratio must be positive", ErrInvalidFrequency)
	}
	return 12 * math.Log2(ratio), nil
}

// NormalizeFrequency normalizes frequency to the nearest reference pitch.
func NormalizeFrequency(frequency float64, reference ...float64) (float64, error) {
	ref := float64(ConcertPitch)
	if len(reference) > 0 {
		ref = reference[0]
	}
	if frequency <= 0 {
		return 0, fmt.Errorf("%w: frequency must be positive", ErrInvalidFrequency)
	}
	nearestA := ref * math.Pow(2, math.Round(12*math.Log2(frequency/ref))/12)
	return nearestA, nil
}