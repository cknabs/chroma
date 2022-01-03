package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Stylus lexer.
var Stylus = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Stylus",
		Aliases:         []string{"stylus"},
		Filenames:       []string{"*.styl"},
		MimeTypes:       []string{"text/x-styl"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/stylus.xml",
))
