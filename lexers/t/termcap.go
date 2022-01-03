package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Termcap lexer.
var Termcap = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Termcap",
		Aliases:   []string{"termcap"},
		Filenames: []string{"termcap", "termcap.src"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/termcap.xml",
))
