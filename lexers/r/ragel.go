package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ragel lexer.
var Ragel = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Ragel",
		Aliases:   []string{"ragel"},
		Filenames: []string{},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/ragel.xml",
))
