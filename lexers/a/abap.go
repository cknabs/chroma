package a

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ABAP lexer.
var Abap = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "ABAP",
		Aliases:         []string{"abap"},
		Filenames:       []string{"*.abap", "*.ABAP"},
		MimeTypes:       []string{"text/x-abap"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/abap.xml",
))
