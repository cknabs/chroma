package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Handlebars lexer.
var Handlebars = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Handlebars",
		Aliases:   []string{"handlebars", "hbs"},
		Filenames: []string{"*.handlebars", "*.hbs"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/handlebars.xml",
))
