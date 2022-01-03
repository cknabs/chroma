package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Turing lexer.
var Turing = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Turing",
		Aliases:   []string{"turing"},
		Filenames: []string{"*.turing", "*.tu"},
		MimeTypes: []string{"text/x-turing"},
	},
	embedded,
	"embedded/turing.xml",
))
