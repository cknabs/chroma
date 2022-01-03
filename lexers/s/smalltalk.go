package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Smalltalk lexer.
var Smalltalk = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Smalltalk",
		Aliases:   []string{"smalltalk", "squeak", "st"},
		Filenames: []string{"*.st"},
		MimeTypes: []string{"text/x-smalltalk"},
	},
	embedded,
	"embedded/smalltalk.xml",
))
