package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ceylon lexer.
var Ceylon = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Ceylon",
		Aliases:   []string{"ceylon"},
		Filenames: []string{"*.ceylon"},
		MimeTypes: []string{"text/x-ceylon"},
		DotAll:    true,
	},
	embedded,
	"embedded/ceylon.xml",
))
