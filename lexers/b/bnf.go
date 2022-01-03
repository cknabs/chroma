package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Bnf lexer.
var Bnf = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "BNF",
		Aliases:   []string{"bnf"},
		Filenames: []string{"*.bnf"},
		MimeTypes: []string{"text/x-bnf"},
	},
	embedded,
	"embedded/bnf.xml",
))
