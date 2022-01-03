package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Prolog lexer.
var Prolog = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Prolog",
		Aliases:   []string{"prolog"},
		Filenames: []string{"*.ecl", "*.prolog", "*.pro", "*.pl"},
		MimeTypes: []string{"text/x-prolog"},
	},
	embedded,
	"embedded/prolog.xml",
))
