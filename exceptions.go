package soundkit

import "errors"

// SoundKitError acts as your base sentinel string error 
// if you want to check for general package errors.
var ErrSoundKit = errors.New("soundkit error")

// Specific, granular error values that your functions can return.
// By Go convention, error variables are prefixed with "Err".
var (
	ErrInvalidNote      = errors.New("invalid note name provided")
	ErrInvalidOctave    = errors.New("octave is out of valid range")
	ErrInvalidFrequency = errors.New("invalid frequency is provided")
	ErrInvalidChord     = errors.New("invalid chord type is provided")
	ErrInvalidScale     = errors.New("invalid scale type is provided")
)