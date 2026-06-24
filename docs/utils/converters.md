# Utils Converters Module Documentation

The `utils/converters` module provides math utilities to convert between frequencies, cents, semitones, and ratios.

## Functions

### `FrequencyToCents(freq1, freq2 float64) (float64, error)`
- **Description**: Calculates the cents difference between two frequency values.
- **Errors**: `ErrInvalidFrequency` if either frequency <= 0.

### `CentsToRatio(cents float64) float64`
- **Description**: Calculates the frequency ratio for a given number of cents.

### `RatioToCents(ratio float64) (float64, error)`
- **Description**: Converts a frequency ratio to cents.
- **Errors**: `ErrInvalidFrequency` if ratio <= 0.

### `SemitonesToRatio(semitones float64) float64`
- **Description**: Converts semitones to a frequency ratio.

### `RatioToSemitones(ratio float64) (float64, error)`
- **Description**: Converts a frequency ratio to semitones.
- **Errors**: `ErrInvalidFrequency` if ratio <= 0.

### `NormalizeFrequency(frequency float64, reference ...float64) (float64, error)`
- **Description**: Normalizes a frequency to the nearest octave reference pitch. The optional `reference` argument defaults to standard ConcertPitch (440Hz).
- **Errors**: `ErrInvalidFrequency` if frequency <= 0.
