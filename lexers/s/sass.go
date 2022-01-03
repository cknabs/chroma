package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Sass lexer.
var Sass = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Sass",
		Aliases:         []string{"sass"},
		Filenames:       []string{"*.sass"},
		MimeTypes:       []string{"text/x-sass"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/sass.xml",
))
