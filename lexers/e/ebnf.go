package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ebnf lexer.
var Ebnf = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "EBNF",
		Aliases:   []string{"ebnf"},
		Filenames: []string{"*.ebnf"},
		MimeTypes: []string{"text/x-ebnf"},
	},
	embedded,
	"embedded/ebnf.xml",
))
