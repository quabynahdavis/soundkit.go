package core

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/quabynahdavis/soundkit.go/utils"
)

// Sentinel errors for core package
var (
	ErrSoundKit      = errors.New("soundkit error")
	ErrInvalidNote   = errors.New("invalid note name provided")
	ErrInvalidOctave = errors.New("octave is out of valid range")
)

// Standard note names and mapping helper constants
var (
	NoteNamesSharp = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	NoteNamesFlat  = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	FlatToSharp    = map[string]string{
		"DB": "C#",
		"EB": "D#",
		"FB": "E",
		"GB": "F#",
		"AB": "G#",
		"BB": "A#",
		"CB": "B",
	}
	ReferenceFrequencies = map[string]float64{
		"C0":  16.35,
		"C#0": 17.32,
		"D0":  18.35,
		"D#0": 19.45,
		"E0":  20.60,
		"F0":  21.83,
		"F#0": 23.12,
		"G0":  24.50,
		"G#0": 25.96,
		"A0":  27.50,
		"A#0": 29.14,
		"B0":  30.87,
		"C4":  261.63,
		"A4":  440.00,
		"C5":  523.25,
	}
)

var noteRegex = regexp.MustCompile(`^([CDEFGAB][B#]?)(-?\d+)$`)

var noteMap = map[string]int{
	"C":  0,
	"C#": 1,
	"DB": 1,
	"D":  2,
	"D#": 3,
	"EB": 3,
	"E":  4,
	"FB": 4,
	"F":  5,
	"F#": 6,
	"GB": 6,
	"G":  7,
	"G#": 8,
	"AB": 8,
	"A":  9,
	"A#": 10,
	"BB": 10,
	"B":  11,
	"CB": 11,
}

// MidiKey converts note name to MIDI note number.
func MidiKey(noteName string) (int, error) {
	cleanNoteName := utils.NormalizeNoteName(noteName)
	match := noteRegex.FindStringSubmatch(cleanNoteName)
	if match == nil {
		return 0, fmt.Errorf("%w: invalid note name format: %s", ErrInvalidNote, noteName)
	}

	note := strings.ToUpper(match[1])
	if len(note) == 2 && note[1] == 'B' {
		if sharp, ok := FlatToSharp[note]; ok {
			note = sharp
		}
	}

	octave, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, fmt.Errorf("%w: unexpected error processing note: %s", ErrInvalidNote, noteName)
	}

	if octave < -1 || octave > 10 {
		return 0, fmt.Errorf("%w: octave out of range (-1 to 10): %d", ErrInvalidOctave, octave)
	}

	noteOffset, ok := noteMap[note]
	if !ok {
		return 0, fmt.Errorf("%w: note doesn't exist: %s", ErrInvalidNote, note)
	}

	midiValue := 12 + (octave * 12) + noteOffset
	if midiValue < 0 || midiValue > 127 {
		return 0, fmt.Errorf("%w: resulting MIDI value out of range: %d", ErrInvalidOctave, midiValue)
	}

	return midiValue, nil
}

// MidiFreq converts note name to frequency.
func MidiFreq(noteName string, roundDigits int, concertPitch float64) (float64, error) {
	pitch, err := MidiKey(noteName)
	if err != nil {
		return 0, err
	}
	freq := math.Pow(2, float64(pitch-69)/12.0) * concertPitch
	if roundDigits >= 0 {
		p := math.Pow(10, float64(roundDigits))
		freq = math.Round(freq*p) / p
	}
	return freq, nil
}

// FreqToMidi converts frequency to MIDI note number.
func FreqToMidi(freq float64, concertPitch float64) (int, error) {
	if freq <= 0 {
		return 0, fmt.Errorf("%w: frequency must be positive", utils.ErrInvalidFrequency)
	}

	midi := 12*math.Log2(freq/concertPitch) + 69
	roundedMidi := int(math.Round(midi))

	if roundedMidi < 0 || roundedMidi > 127 {
		return 0, fmt.Errorf("%w: frequency %fHz results in out-of-range MIDI note: %d", utils.ErrInvalidFrequency, freq, roundedMidi)
	}

	return roundedMidi, nil
}

// MidiToNoteName converts MIDI number to note name.
func MidiToNoteName(midiNumber int, useSharps bool) (string, error) {
	if midiNumber < 0 || midiNumber > 127 {
		return "", fmt.Errorf("%w: MIDI number out of range (0-127): %d", ErrInvalidNote, midiNumber)
	}

	var notes []string
	if useSharps {
		notes = NoteNamesSharp
	} else {
		notes = NoteNamesFlat
	}

	octave := (midiNumber / 12) - 1
	note := notes[midiNumber%12]
	return fmt.Sprintf("%s%d", note, octave), nil
}

// IsValidMidiRange checks if note is within standard MIDI range (0-127).
func IsValidMidiRange(noteName string) bool {
	midiVal, err := MidiKey(noteName)
	if err != nil {
		return false
	}
	return midiVal >= 0 && midiVal <= 127
}

// NotesToFrequencies converts list of notes to frequencies.
func NotesToFrequencies(noteList []string, roundDigits int, concertPitch float64) []any {
	results := make([]any, 0, len(noteList))
	for _, note := range noteList {
		freq, err := MidiFreq(note, roundDigits, concertPitch)
		if err != nil {
			results = append(results, err.Error())
		} else {
			results = append(results, freq)
		}
	}
	return results
}

// NotesToMidi converts list of notes to MIDI numbers.
func NotesToMidi(noteList []string) []any {
	results := make([]any, 0, len(noteList))
	for _, note := range noteList {
		midi, err := MidiKey(note)
		if err != nil {
			results = append(results, err.Error())
		} else {
			results = append(results, midi)
		}
	}
	return results
}
