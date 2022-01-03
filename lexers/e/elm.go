package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Elm lexer.
var Elm = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Elm",
		Aliases:   []string{"elm"},
		Filenames: []string{"*.elm"},
		MimeTypes: []string{"text/x-elm"},
	},
	embedded,
	"embedded/elm.xml",
))
