package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Tcsh lexer.
var Tcsh = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Tcsh",
		Aliases:   []string{"tcsh", "csh"},
		Filenames: []string{"*.tcsh", "*.csh"},
		MimeTypes: []string{"application/x-csh"},
	},
	embedded,
	"embedded/tcsh.xml",
))
