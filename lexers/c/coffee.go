package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Coffeescript lexer.
var Coffeescript = internal.Register(MustNewXMLLexer(
	&Config{
		Name:         "CoffeeScript",
		Aliases:      []string{"coffee-script", "coffeescript", "coffee"},
		Filenames:    []string{"*.coffee"},
		MimeTypes:    []string{"text/coffeescript"},
		NotMultiline: true,
		DotAll:       true,
	},
	embedded,
	"embedded/coffeescript.xml",
))
