package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Turtle lexer.
var Turtle = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Turtle",
		Aliases:         []string{"turtle"},
		Filenames:       []string{"*.ttl"},
		MimeTypes:       []string{"text/turtle", "application/x-turtle"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/turtle.xml",
))
