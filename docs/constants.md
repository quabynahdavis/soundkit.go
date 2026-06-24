# Constants Documentation

The `soundkit` package exports key musical constants and mapping references.

## Constants

### `ConcertPitch`
- **Type**: `int` (440)
- **Description**: Standard tuning frequency (A4 = 440Hz).

## Variables

### `MidiRange`
- **Type**: `[2]int{0, 127}`
- **Description**: Valid range of MIDI note numbers.

### `ReferenceFrequencies`
- **Type**: `map[string]float64`
- **Description**: Map of reference note names to their exact frequencies (e.g. C0 -> 16.35, A4 -> 440.00).

### `NoteNamesSharp`
- **Type**: `[]string`
- **Description**: Chromatic scale note names using sharps: `["C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"]`.

### `NoteNamesFlat`
- **Type**: `[]string`
- **Description**: Chromatic scale note names using flats: `["C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"]`.

### `FlatToSharp`
- **Type**: `map[string]string`
- **Description**: Mapping from uppercase flat note names to their sharp equivalents (e.g., `DB` -> `C#`).
