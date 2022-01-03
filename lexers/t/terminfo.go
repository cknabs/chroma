package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Terminfo lexer.
var Terminfo = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Terminfo",
		Aliases:   []string{"terminfo"},
		Filenames: []string{"terminfo", "terminfo.src"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/terminfo.xml",
))
