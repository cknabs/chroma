package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Promql lexer.
var Promql = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "PromQL",
		Aliases:   []string{"promql"},
		Filenames: []string{"*.promql"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/promql.xml",
))
