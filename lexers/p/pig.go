package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Pig lexer.
var Pig = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Pig",
		Aliases:         []string{"pig"},
		Filenames:       []string{"*.pig"},
		MimeTypes:       []string{"text/x-pig"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/pig.xml",
))
