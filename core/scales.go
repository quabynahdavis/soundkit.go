package core

import (
	"errors"
	"fmt"

	"github.com/quabynahdavis/soundkit.go/utils"
)

// ErrInvalidScale is raised when an invalid scale type is provided.
var ErrInvalidScale = errors.New("invalid scale type is provided")

// ScaleTypes defines common musical scales and their intervals.
var ScaleTypes = map[string][]int{
	"major":            {0, 2, 4, 5, 7, 9, 11},
	"minor":            {0, 2, 3, 5, 7, 8, 10},
	"natural_minor":    {0, 2, 3, 5, 7, 8, 10},
	"harmonic_minor":   {0, 2, 3, 5, 7, 8, 11},
	"melodic_minor":    {0, 2, 3, 5, 7, 9, 11},
	"pentatonic_major": {0, 2, 4, 7, 9},
	"pentatonic_minor": {0, 3, 5, 7, 10},
	"blues":            {0, 3, 5, 6, 7, 10},
	"dorian":           {0, 2, 3, 5, 7, 9, 10},
	"phrygian":         {0, 1, 3, 5, 7, 8, 10},
	"lydian":           {0, 2, 4, 6, 7, 9, 11},
	"mixolydian":       {0, 2, 4, 5, 7, 9, 10},
	"locrian":          {0, 1, 3, 5, 6, 8, 10},
	"whole_tone":       {0, 2, 4, 6, 8, 10},
}

// GetScaleNotes gets MIDI notes for common scales.
func GetScaleNotes(scaleRoot string, scaleType string, octave int, numOctaves int) ([]int, error) {
	intervals, ok := ScaleTypes[scaleType]
	if !ok {
		return nil, fmt.Errorf("%w: unknown scale type: %s", ErrInvalidScale, scaleType)
	}

	rootMidi, err := MidiKey(fmt.Sprintf("%s%d", scaleRoot, octave))
	if err != nil {
		return nil, fmt.Errorf("%w: invalid scale root: %s", ErrInvalidScale, scaleRoot)
	}

	var notes []int
	for octaveOffset := 0; octaveOffset < numOctaves; octaveOffset++ {
		for _, interval := range intervals {
			noteMidi := rootMidi + interval + (12 * octaveOffset)
			if noteMidi <= 127 { // Stay within MIDI range
				notes = append(notes, noteMidi)
			}
		}
	}
	return notes, nil
}

// GetScaleFrequencies gets frequencies for a scale.
func GetScaleFrequencies(scaleRoot string, scaleType string, octave int, numOctaves int, roundDigits int) ([]float64, error) {
	midiNotes, err := GetScaleNotes(scaleRoot, scaleType, octave, numOctaves)
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

// GetScaleNames gets list of available scale types.
func GetScaleNames() []string {
	names := make([]string, 0, len(ScaleTypes))
	for name := range ScaleTypes {
		names = append(names, name)
	}
	return names
}
