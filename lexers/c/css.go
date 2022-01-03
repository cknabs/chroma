package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CSS lexer.
var CSS = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "CSS",
		Aliases:   []string{"css"},
		Filenames: []string{"*.css"},
		MimeTypes: []string{"text/css"},
	},
	embedded,
	"embedded/css.xml",
))
