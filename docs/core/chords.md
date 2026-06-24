# Core Chords Module Documentation

The `core/chords` module generates chord intervals, support for inversions, and chord frequencies.

## Chord Types

Supported chord types:
- `maj` / `major`, `min` / `minor`
- `dim` / `diminished`, `aug` / `augmented`
- `7` / `dominant7`, `maj7` / `major7`, `min7` / `minor7`
- `dim7` / `diminished7`, `half_dim7` / `m7b5`
- `sus2`, `sus4`, `9`, `maj9`

## Functions

### `GetChordNotes(chordRoot string, chordType string, octave int, inversion int) ([]int, error)`
- **Description**: Generates MIDI notes for a given chord starting at `chordRoot` in the specified `octave` with optional `inversion` offset (e.g. 1st, 2nd inversion).
- **Errors**: `ErrInvalidChord`, `ErrInvalidNote`.

### `GetChordFrequencies(chordRoot string, chordType string, octave int, roundDigits int, inversion int) ([]float64, error)`
- **Description**: Returns frequency values for the generated chord.
- **Errors**: `ErrInvalidChord`, `ErrInvalidNote`.

### `GetChordNames() []string`
- **Description**: Returns all registered chord types.
