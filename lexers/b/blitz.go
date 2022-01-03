package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Blitzbasic lexer.
var Blitzbasic = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "BlitzBasic",
		Aliases:         []string{"blitzbasic", "b3d", "bplus"},
		Filenames:       []string{"*.bb", "*.decls"},
		MimeTypes:       []string{"text/x-bb"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/blitzbasic.xml",
))
