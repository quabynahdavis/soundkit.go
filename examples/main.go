package main

import (
	"fmt"
	"log"

	"github.com/quabynahdavis/soundkit.go"
)

func main() {
	fmt.Println("=== SoundKit Go Examples ===\n")

	// 1. Basic Note Conversion
	fmt.Println("1. Basic Note Conversion:")
	notes := []string{"C4", "A4", "Gb3", "D#5"}
	for _, note := range notes {
		midi, err := soundkit.MidiKey(note)
		if err != nil {
			log.Fatalf("Error getting MIDI key: %v", err)
		}

		freq, err := soundkit.MidiFreq(note, 2, float64(soundkit.ConcertPitch))
		if err != nil {
			log.Fatalf("Error getting frequency: %v", err)
		}
		fmt.Printf("  %s -> MIDI: %d, Frequency: %.2fHz\n", note, midi, freq)
	}

	// 2. Chord Examples
	fmt.Println("\n2. Chord Examples:")
	chords := []struct {
		root      string
		chordType string
	}{
		{"C", "maj"},
		{"D", "min7"},
		{"G", "7"},
	}
	for _, c := range chords {
		chordNotes, err := soundkit.GetChordNotes(c.root, c.chordType, 4, 0)
		if err != nil {
			log.Fatalf("Error generating chord: %v", err)
		}

		chordFreqs, err := soundkit.GetChordFrequencies(c.root, c.chordType, 4, 2, 0)
		if err != nil {
			log.Fatalf("Error generating chord frequencies: %v", err)
		}
		fmt.Printf("  %s%s: %v -> %vHz\n", c.root, c.chordType, chordNotes, chordFreqs)
	}

	// 3. Scale Examples
	fmt.Println("\n3. Scale Examples:")
	scales := []struct {
		root      string
		scaleType string
	}{
		{"C", "major"},
		{"A", "minor"},
	}
	for _, s := range scales {
		scaleNotes, err := soundkit.GetScaleNotes(s.root, s.scaleType, 4, 1)
		if err != nil {
			log.Fatalf("Error generating scale: %v", err)
		}
		fmt.Printf("  %s %s: %v\n", s.root, s.scaleType, scaleNotes)
	}

	// 4. Reverse Conversion
	fmt.Println("\n4. Reverse Conversion:")
	midiNote := 60
	noteName, err := soundkit.MidiToNoteName(midiNote, true)
	if err != nil {
		log.Fatalf("Error converting MIDI to note name: %v", err)
	}
	freq, err := soundkit.MidiFreq(noteName, 2, float64(soundkit.ConcertPitch))
	if err != nil {
		log.Fatalf("Error getting frequency: %v", err)
	}
	fmt.Printf("  MIDI %d -> %s -> %.2fHz\n", midiNote, noteName, freq)
}
