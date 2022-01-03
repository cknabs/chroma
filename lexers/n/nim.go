package n

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Nim lexer.
var Nim = internal.Register(MustNewXMLLexer(
	&Config{
		Name:            "Nim",
		Aliases:         []string{"nim", "nimrod"},
		Filenames:       []string{"*.nim", "*.nimrod"},
		MimeTypes:       []string{"text/x-nim"},
		CaseInsensitive: true,
	},
	embedded,
	"embedded/nim.xml",
))
