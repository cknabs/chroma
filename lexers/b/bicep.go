package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Bicep lexer.
var Bicep = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Bicep",
		Aliases:   []string{"bicep"},
		Filenames: []string{"*.bicep"},
	},
	embedded,
	"embedded/bicep.xml",
))
