package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ANTLR lexer.
var ANTLR = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "ANTLR",
		Aliases:   []string{"antlr"},
		Filenames: []string{},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/antlr.xml",
))
