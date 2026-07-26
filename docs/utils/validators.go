# Utils Validators Module Documentation

The `utils/validators` module provides format and range validation.

## Functions

### `ValidateNoteName(noteName string) bool`
- **Description**: Checks if a string matches standard note formatting rules (e.g. `C4`, `A#9`, `Eb-1`). Replaces spaces and executes a regex match.

### `ValidateMidiRange(midiNumber int) bool`
- **Description**: Returns `true` if the MIDI note number is within `[0, 127]`.

### `ValidateFrequency(frequency float64) bool`
- **Description**: Checks if the frequency is positive and falls within the human hearing range (`(0, 20000]`).

### `ValidateOctave(octave int) bool`
- **Description**: Checks if the octave number is within standard musical range (`[-1, 10]`).

### `NormalizeNoteName(noteName string) string`
- **Description**: Standardizes formatting by converting string to uppercase, removing spaces and hyphens, and converting unicode flat/sharp symbols (`♭`, `♯`) to standard characters (`B`, `#`).
