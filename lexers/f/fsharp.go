package f

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Fsharp lexer.
var Fsharp = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "FSharp",
		Aliases:   []string{"fsharp"},
		Filenames: []string{"*.fs", "*.fsi"},
		MimeTypes: []string{"text/x-fsharp"},
	},
	embedded,
	"embedded/fsharp.xml",
))
