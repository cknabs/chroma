package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Hexdump lexer.
var Hexdump = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Hexdump",
		Aliases:   []string{"hexdump"},
		Filenames: []string{},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/hexdump.xml",
))
