package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Abnf lexer.
var Abnf = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "ABNF",
		Aliases:   []string{"abnf"},
		Filenames: []string{"*.abnf"},
		MimeTypes: []string{"text/x-abnf"},
	},
	embedded,
	"embedded/abnf.xml",
))
