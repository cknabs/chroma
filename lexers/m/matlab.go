package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Matlab lexer.
var Matlab = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Matlab",
		Aliases:   []string{"matlab"},
		Filenames: []string{"*.m"},
		MimeTypes: []string{"text/matlab"},
	},
	embedded,
	"embedded/matlab.xml",
))
