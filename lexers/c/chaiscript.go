package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Chaiscript lexer.
var Chaiscript = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "ChaiScript",
		Aliases:   []string{"chai", "chaiscript"},
		Filenames: []string{"*.chai"},
		MimeTypes: []string{"text/x-chaiscript", "application/x-chaiscript"},
		DotAll:    true,
	},
	embedded,
	"embedded/chaiscript.xml",
))
