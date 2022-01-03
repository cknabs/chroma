package m

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Modula-2 lexer.
var Modula2 = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Modula-2",
		Aliases:   []string{"modula2", "m2"},
		Filenames: []string{"*.def", "*.mod"},
		MimeTypes: []string{"text/x-modula2"},
		DotAll:    true,
	},
	embedded,
	"embedded/modula-2.xml",
))
