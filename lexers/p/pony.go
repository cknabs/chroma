package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pony lexer.
var Pony = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Pony",
		Aliases:   []string{"pony"},
		Filenames: []string{"*.pony"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/pony.xml",
))
