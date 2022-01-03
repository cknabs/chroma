package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Postscript lexer.
var Postscript = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "PostScript",
		Aliases:   []string{"postscript", "postscr"},
		Filenames: []string{"*.ps", "*.eps"},
		MimeTypes: []string{"application/postscript"},
	},
	embedded,
	"embedded/postscript.xml",
))
