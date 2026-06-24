package soundkit

import (
	"github.com/quabynahdavis/soundkit.go/core"
	"github.com/quabynahdavis/soundkit.go/utils"
)

// ConcertPitch is the standard tuning frequency (440Hz for Key A).
const ConcertPitch = utils.ConcertPitch

var (
	// MidiRange defines the valid range of MIDI numbers (0 to 127).
	MidiRange = utils.MidiRange

	// ReferenceFrequencies is a map of reference frequencies for key pitches.
	ReferenceFrequencies = core.ReferenceFrequencies

	// NoteNamesSharp is standard note names with sharp representation.
	NoteNamesSharp = core.NoteNamesSharp

	// NoteNamesFlat is standard note names with flat representation.
	NoteNamesFlat = core.NoteNamesFlat

	// FlatToSharp maps flat representation note names to sharp names.
	FlatToSharp = core.FlatToSharp
)
