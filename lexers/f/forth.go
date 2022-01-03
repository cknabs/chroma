package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Forth lexer.
var Forth = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Forth",
		Aliases:         []string{"forth"},
		Filenames:       []string{"*.frt", "*.fth", "*.fs"},
		MimeTypes:       []string{"application/x-forth"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/forth.xml",
))
