package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Groff lexer.
var Groff = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Groff",
		Aliases:   []string{"groff", "nroff", "man"},
		Filenames: []string{"*.[1-9]", "*.1p", "*.3pm", "*.man"},
		MimeTypes: []string{"application/x-troff", "text/troff"},
	},
	embedded,
	"embedded/groff.xml",
))
