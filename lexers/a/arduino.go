package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Arduino lexer.
var Arduino = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Arduino",
		Aliases:   []string{"arduino"},
		Filenames: []string{"*.ino"},
		MimeTypes: []string{"text/x-arduino"},
		EnsureNL:  true,
	},
	embedded,
	"embedded/arduino.xml",
))
