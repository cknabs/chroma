package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Crystal lexer.
var Crystal = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Crystal",
		Aliases:   []string{"cr", "crystal"},
		Filenames: []string{"*.cr"},
		MimeTypes: []string{"text/x-crystal"},
		DotAll:    true,
	},
	embedded,
	"embedded/crystal.xml",
))
