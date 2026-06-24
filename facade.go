package soundkit

import (
	"github.com/quabynahdavis/soundkit.go/core"
	"github.com/quabynahdavis/soundkit.go/utils"
)

// Validators & Converters (from utils)

// ValidateNoteName validates if a string is a properly formatted note name.
func ValidateNoteName(noteName string) bool {
	return utils.ValidateNoteName(noteName)
}

// ValidateMidiRange validates if a number is within the MIDI range.
func ValidateMidiRange(midiNumber int) bool {
	return utils.ValidateMidiRange(midiNumber)
}

// ValidateFrequency validates if a frequency is positive and reasonable (within human hearing range).
func ValidateFrequency(frequency float64) bool {
	return utils.ValidateFrequency(frequency)
}

// ValidateOctave validates if an octave is within reasonable range (-1 to 10).
func ValidateOctave(octave int) bool {
	return utils.ValidateOctave(octave)
}

// NormalizeNoteName converts various note formats to standard format.
func NormalizeNoteName(noteName string) string {
	return utils.NormalizeNoteName(noteName)
}

// FrequencyToCents converts frequency ratio to cents.
func FrequencyToCents(freq1, freq2 float64) (float64, error) {
	return utils.FrequencyToCents(freq1, freq2)
}

// CentsToRatio converts cents to frequency ratio.
func CentsToRatio(cents float64) float64 {
	return utils.CentsToRatio(cents)
}

// RatioToCents converts frequency ratio to cents.
func RatioToCents(ratio float64) (float64, error) {
	return utils.RatioToCents(ratio)
}

// SemitonesToRatio converts semitones to frequency ratio.
func SemitonesToRatio(semitones float64) float64 {
	return utils.SemitonesToRatio(semitones)
}

// RatioToSemitones converts frequency ratio to semitones.
func RatioToSemitones(ratio float64) (float64, error) {
	return utils.RatioToSemitones(ratio)
}

// NormalizeFrequency normalizes frequency to the nearest reference pitch.
func NormalizeFrequency(frequency float64, reference ...float64) (float64, error) {
	return utils.NormalizeFrequency(frequency, reference...)
}

// Core notes, scales, and chords (from core)

// MidiKey converts note name to MIDI note number.
func MidiKey(noteName string) (int, error) {
	return core.MidiKey(noteName)
}

// MidiFreq converts note name to frequency.
func MidiFreq(noteName string, roundDigits int, concertPitch float64) (float64, error) {
	return core.MidiFreq(noteName, roundDigits, concertPitch)
}

// FreqToMidi converts frequency to MIDI note number.
func FreqToMidi(freq float64, concertPitch float64) (int, error) {
	return core.FreqToMidi(freq, concertPitch)
}

// MidiToNoteName converts MIDI number to note name.
func MidiToNoteName(midiNumber int, useSharps bool) (string, error) {
	return core.MidiToNoteName(midiNumber, useSharps)
}

// IsValidMidiRange checks if note is within standard MIDI range (0-127).
func IsValidMidiRange(noteName string) bool {
	return core.IsValidMidiRange(noteName)
}

// NotesToFrequencies converts list of notes to frequencies.
func NotesToFrequencies(noteList []string, roundDigits int, concertPitch float64) []any {
	return core.NotesToFrequencies(noteList, roundDigits, concertPitch)
}

// NotesToMidi converts list of notes to MIDI numbers.
func NotesToMidi(noteList []string) []any {
	return core.NotesToMidi(noteList)
}

// GetScaleNotes gets MIDI notes for common scales.
func GetScaleNotes(scaleRoot string, scaleType string, octave int, numOctaves int) ([]int, error) {
	return core.GetScaleNotes(scaleRoot, scaleType, octave, numOctaves)
}

// GetScaleFrequencies gets frequencies for a scale.
func GetScaleFrequencies(scaleRoot string, scaleType string, octave int, numOctaves int, roundDigits int) ([]float64, error) {
	return core.GetScaleFrequencies(scaleRoot, scaleType, octave, numOctaves, roundDigits)
}

// GetScaleNames gets list of available scale types.
func GetScaleNames() []string {
	return core.GetScaleNames()
}

// GetChordNotes gets MIDI notes for common chords with optional inversion.
func GetChordNotes(chordRoot string, chordType string, octave int, inversion int) ([]int, error) {
	return core.GetChordNotes(chordRoot, chordType, octave, inversion)
}

// GetChordFrequencies gets frequencies for a chord.
func GetChordFrequencies(chordRoot string, chordType string, octave int, roundDigits int, inversion int) ([]float64, error) {
	return core.GetChordFrequencies(chordRoot, chordType, octave, roundDigits, inversion)
}

// GetChordNames gets list of available chord types.
func GetChordNames() []string {
	return core.GetChordNames()
}
