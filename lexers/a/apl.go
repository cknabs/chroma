package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Apl lexer.
var Apl = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "APL",
		Aliases:   []string{"apl"},
		Filenames: []string{"*.apl"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/apl.xml",
))
