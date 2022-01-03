package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// J lexer.
var J = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "J",
		Aliases:   []string{"j"},
		Filenames: []string{"*.ijs"},
		MimeTypes: []string{"text/x-j"},
	},
	embedded,
	"embedded/j.xml",
))
