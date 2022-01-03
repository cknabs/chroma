package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Coq lexer.
var Coq = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Coq",
		Aliases:   []string{"coq"},
		Filenames: []string{"*.v"},
		MimeTypes: []string{"text/x-coq"},
	},
	embedded,
	"embedded/coq.xml",
))
