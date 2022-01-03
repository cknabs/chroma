package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Julia lexer.
var Julia = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Julia",
		Aliases:   []string{"julia", "jl"},
		Filenames: []string{"*.jl"},
		MimeTypes: []string{"text/x-julia", "application/x-julia"},
	},
	embedded,
	"embedded/julia.xml",
))
