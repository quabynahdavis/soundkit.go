package soundkit

import (
	"github.com/quabynahdavis/soundkit.go/core"
	"github.com/quabynahdavis/soundkit.go/utils"
)

var (
	// ErrSoundKit is the base sentinel error for all SoundKit errors.
	ErrSoundKit = core.ErrSoundKit

	// ErrInvalidNote is raised when an invalid note name is provided.
	ErrInvalidNote = core.ErrInvalidNote

	// ErrInvalidOctave is raised when the octave is out of valid range.
	ErrInvalidOctave = core.ErrInvalidOctave

	// ErrInvalidFrequency is raised when an invalid frequency is provided.
	ErrInvalidFrequency = utils.ErrInvalidFrequency

	// ErrInvalidChord is raised when an invalid chord type is provided.
	ErrInvalidChord = core.ErrInvalidChord

	// ErrInvalidScale is raised when an invalid scale type is provided.
	ErrInvalidScale = core.ErrInvalidScale
)