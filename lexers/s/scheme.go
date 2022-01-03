package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// nolint

// Scheme lexer.
var SchemeLang = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Scheme",
		Aliases:   []string{"scheme", "scm"},
		Filenames: []string{"*.scm", "*.ss"},
		MimeTypes: []string{"text/x-scheme", "application/x-scheme"},
	},
	embedded,
	"embedded/scheme.xml",
))
