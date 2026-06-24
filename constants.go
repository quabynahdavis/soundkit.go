package soundkit

// Constants are strictly for basic scalar types (numbers, strings, booleans)
const CONCERT_PITCH = 440 // Key A
var MIDI_RANGE = [2]int{0, 127}

var REFERENCE_FREQUENCIES = map[string]float32{
	"C0":   16.35,
	"C#0":  17.32,
	"D0":   18.35,
	"D#0":  19.45,
	"E0":   20.60,
	"F0":   21.83,
	"F#0":  23.12,
	"G0":   24.50,
	"G#0":  25.96,
	"A0":   27.50,
	"A#0":  29.14,
	"B0":   30.87,
	"C4":   261.63,
	"A4":   440.00,
	"C5":   523.25,
}

// Standard note names
var NOTE_NAMES_SHARP = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var NOTE_NAMES_FLAT =[]string {"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var FLAT_TO_SHARP=  map[string] string {
    "DB": "C#",
    "EB": "D#",
    "FB": "E",
    "GB": "F#",
    "AB": "G#",
    "BB": "A#",
    "CB": "B",
}
