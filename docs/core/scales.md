# Core Scales Module Documentation

The `core/scales` module generates note intervals and frequencies for musical scales.

## Scale Types

Supported scale types:
- `major`, `minor`, `natural_minor`, `harmonic_minor`, `melodic_minor`
- `pentatonic_major`, `pentatonic_minor`
- `blues`
- `dorian`, `phrygian`, `lydian`, `mixolydian`, `locrian`
- `whole_tone`

## Functions

### `GetScaleNotes(scaleRoot string, scaleType string, octave int, numOctaves int) ([]int, error)`
- **Description**: Generates MIDI note numbers for a given scale, starting from `scaleRoot` (e.g. "C") in the specified `octave`, spanning `numOctaves`.
- **Errors**: `ErrInvalidScale`, `ErrInvalidNote`.

### `GetScaleFrequencies(scaleRoot string, scaleType string, octave int, numOctaves int, roundDigits int) ([]float64, error)`
- **Description**: Generates frequency floats for a given scale.
- **Errors**: `ErrInvalidScale`, `ErrInvalidNote`.

### `GetScaleNames() []string`
- **Description**: Returns all registered scale types.
