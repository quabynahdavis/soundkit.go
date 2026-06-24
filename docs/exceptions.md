# Exceptions & Sentinel Errors Documentation

SoundKit uses Go idiomatic sentinel errors for error handling.

## Sentinel Errors

### `ErrSoundKit`
- **Description**: Base error for all SoundKit errors.

### `ErrInvalidNote`
- **Description**: Raised when an invalid note name is provided.

### `ErrInvalidOctave`
- **Description**: Raised when the octave is out of valid range (-1 to 10).

### `ErrInvalidFrequency`
- **Description**: Raised when an invalid frequency is provided (e.g. <= 0).

### `ErrInvalidChord`
- **Description**: Raised when an invalid chord type is provided.

### `ErrInvalidScale`
- **Description**: Raised when an invalid scale type is provided.

## Usage Example

```go
package main

import (
	"errors"
	"fmt"
	
	"github.com/quabynahdavis/soundkit.go"
)

func main() {
	_, err := soundkit.MidiKey("H4")
	if errors.Is(err, soundkit.ErrInvalidNote) {
		fmt.Println("Error: Invalid note name provided!")
	}
}
```
