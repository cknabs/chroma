package w

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// WDTE lexer.
var WDTE = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "WDTE",
		Filenames: []string{"*.wdte"},
	},
	embedded,
	"embedded/wdte.xml",
))
