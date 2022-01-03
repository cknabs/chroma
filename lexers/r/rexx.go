package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Rexx lexer.
var Rexx = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Rexx",
		Aliases:         []string{"rexx", "arexx"},
		Filenames:       []string{"*.rexx", "*.rex", "*.rx", "*.arexx"},
		MimeTypes:       []string{"text/x-rexx"},
		NotMultiline:    true,
		CaseInsensitive: true,
	},
	embedded,
	"embedded/rexx.xml",
))
