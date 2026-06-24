package utils

import (
	"regexp"
	"strings"
)

// MidiRange defines the valid range of MIDI numbers (0 to 127).
var MidiRange = [2]int{0, 127}

// noteRegex matches a valid note name (e.g., C4, C#4, Db-1, Gb9).
var noteRegex = regexp.MustCompile(`^[CDEFGAB][B#]?-?\d+$`)

// ValidateNoteName validates if a string is a properly formatted note name.
func ValidateNoteName(noteName string) bool {
	// Replaces spaces and matches regex
	clean := strings.ToUpper(strings.ReplaceAll(noteName, " ", ""))
	return noteRegex.MatchString(clean)
}

// ValidateMidiRange validates if a number is within the MIDI range.
func ValidateMidiRange(midiNumber int) bool {
	return midiNumber >= MidiRange[0] && midiNumber <= MidiRange[1]
}

// ValidateFrequency validates if a frequency is positive and reasonable (within human hearing range).
func ValidateFrequency(frequency float64) bool {
	return frequency > 0 && frequency <= 20000
}

// ValidateOctave validates if an octave is within reasonable range (-1 to 10).
func ValidateOctave(octave int) bool {
	return octave >= -1 && octave <= 10
}

// NormalizeNoteName converts various note formats to standard format.
func NormalizeNoteName(noteName string) string {
	res := strings.ToUpper(noteName)
	res = strings.ReplaceAll(res, "-", "")
	res = strings.ReplaceAll(res, " ", "")
	res = strings.ReplaceAll(res, "♭", "B")
	res = strings.ReplaceAll(res, "♯", "#")
	return res
}
