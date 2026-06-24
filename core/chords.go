package core

import (
	"errors"
	"fmt"

	"github.com/quabynahdavis/soundkit.go/utils"
)

// ErrInvalidChord is raised when an invalid chord type is provided.
var ErrInvalidChord = errors.New("invalid chord type is provided")

// ChordTypes defines common musical chords and their intervals.
var ChordTypes = map[string][]int{
	"maj":         {0, 4, 7},
	"major":       {0, 4, 7},
	"min":         {0, 3, 7},
	"minor":       {0, 3, 7},
	"dim":         {0, 3, 6},
	"diminished":  {0, 3, 6},
	"aug":         {0, 4, 8},
	"augmented":   {0, 4, 8},
	"7":           {0, 4, 7, 10},
	"dominant7":   {0, 4, 7, 10},
	"maj7":        {0, 4, 7, 11},
	"major7":      {0, 4, 7, 11},
	"min7":        {0, 3, 7, 10},
	"minor7":      {0, 3, 7, 10},
	"dim7":        {0, 3, 6, 9},
	"diminished7": {0, 3, 6, 9},
	"half_dim7":   {0, 3, 6, 10},
	"m7b5":        {0, 3, 6, 10},
	"sus2":        {0, 2, 7},
	"sus4":        {0, 5, 7},
	"9":           {0, 4, 7, 10, 14},
	"maj9":        {0, 4, 7, 11, 14},
}

// GetChordNotes gets MIDI notes for common chords with optional inversion.
func GetChordNotes(chordRoot string, chordType string, octave int, inversion int) ([]int, error) {
	intervals, ok := ChordTypes[chordType]
	if !ok {
		return nil, fmt.Errorf("%w: unknown chord type: %s", ErrInvalidChord, chordType)
	}

	rootMidi, err := MidiKey(fmt.Sprintf("%s%d", chordRoot, octave))
	if err != nil {
		return nil, fmt.Errorf("%w: invalid chord root: %s", ErrInvalidChord, chordRoot)
	}

	notes := make([]int, len(intervals))
	for i, interval := range intervals {
		notes[i] = rootMidi + interval
	}

	// Apply inversion
	if inversion > 0 {
		for k := 0; k < inversion; k++ {
			if len(notes) > 0 {
				note := notes[0]
				notes = notes[1:]
				notes = append(notes, note+12)
			}
		}
	}
	return notes, nil
}

// GetChordFrequencies gets frequencies for a chord.
func GetChordFrequencies(chordRoot string, chordType string, octave int, roundDigits int, inversion int) ([]float64, error) {
	midiNotes, err := GetChordNotes(chordRoot, chordType, octave, inversion)
	if err != nil {
		return nil, err
	}

	freqs := make([]float64, len(midiNotes))
	for i, note := range midiNotes {
		noteName, err := MidiToNoteName(note, true)
		if err != nil {
			return nil, err
		}
		freq, err := MidiFreq(noteName, roundDigits, float64(utils.ConcertPitch))
		if err != nil {
			return nil, err
		}
		freqs[i] = freq
	}
	return freqs, nil
}

// GetChordNames gets list of available chord types.
func GetChordNames() []string {
	names := make([]string, 0, len(ChordTypes))
	for name := range ChordTypes {
		names = append(names, name)
	}
	return names
}
