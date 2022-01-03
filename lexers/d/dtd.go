package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Dtd lexer.
var Dtd = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "DTD",
		Aliases:   []string{"dtd"},
		Filenames: []string{"*.dtd"},
		MimeTypes: []string{"application/xml-dtd"},
		DotAll:    true,
	},
	embedded,
	"embedded/dtd.xml",
))
