package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ballerina lexer.
var Ballerina = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Ballerina",
		Aliases:   []string{"ballerina"},
		Filenames: []string{"*.bal"},
		MimeTypes: []string{"text/x-ballerina"},
		DotAll:    true,
	},
	embedded,
	"embedded/ballerina.xml",
))
