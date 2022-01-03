package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Povray lexer.
var Povray = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "POVRay",
		Aliases:   []string{"pov"},
		Filenames: []string{"*.pov", "*.inc"},
		MimeTypes: []string{"text/x-povray"},
	},
	embedded,
	"embedded/povray.xml",
))
