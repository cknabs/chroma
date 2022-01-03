package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Octave lexer.
var Octave = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Octave",
		Aliases:   []string{"octave"},
		Filenames: []string{"*.m"},
		MimeTypes: []string{"text/octave"},
	},
	embedded,
	"embedded/octave.xml",
))
