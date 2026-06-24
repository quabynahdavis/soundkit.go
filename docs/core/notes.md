# Core Notes Module Documentation

The `core/notes` module contains core logic to handle note parsing, conversion between note names, MIDI keys, and frequencies.

## Functions

### `MidiKey(noteName string) (int, error)`
- **Description**: Converts a note name (e.g. "C4", "A#4", "Gb3") to a MIDI note number.
- **Errors**: `ErrInvalidNote`, `ErrInvalidOctave`.

### `MidiFreq(noteName string, roundDigits int, concertPitch float64) (float64, error)`
- **Description**: Converts a note name to a frequency. Replaces sharps/flats, calculates pitch offsets, and applies optional decimal rounding if `roundDigits` is non-negative.
- **Errors**: `ErrInvalidNote`, `ErrInvalidOctave`.

### `FreqToMidi(freq float64, concertPitch float64) (int, error)`
- **Description**: Converts a frequency value to the closest MIDI note number.
- **Errors**: `ErrInvalidFrequency`.

### `MidiToNoteName(midiNumber int, useSharps bool) (string, error)`
- **Description**: Converts a MIDI note number to its standard note name format.
- **Errors**: `ErrInvalidNote`.

### `IsValidMidiRange(noteName string) bool`
- **Description**: Returns `true` if the note name represents a valid MIDI key (range 0 to 127).

### `NotesToFrequencies(noteList []string, roundDigits int, concertPitch float64) []any`
- **Description**: Batch converts a list of note names to a list of frequency floats or string error messages (union type represented as `[]any`).

### `NotesToMidi(noteList []string) []any`
- **Description**: Batch converts a list of note names to a list of MIDI note numbers or string error messages.
