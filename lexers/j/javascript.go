package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Javascript lexer.
var Javascript = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "JavaScript",
		Aliases:   []string{"js", "javascript"},
		Filenames: []string{"*.js", "*.jsm", "*.mjs"},
		MimeTypes: []string{"application/javascript", "application/x-javascript", "text/x-javascript", "text/javascript"},
		DotAll:    true,
		EnsureNL:  true,
	},
	embedded,
	"embedded/javascript.xml",
))
