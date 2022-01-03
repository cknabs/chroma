package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Snobol lexer.
var Snobol = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Snobol",
		Aliases:   []string{"snobol"},
		Filenames: []string{"*.snobol"},
		MimeTypes: []string{"text/x-snobol"},
	},
	embedded,
	"embedded/snobol.xml",
))
