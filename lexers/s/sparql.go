package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Sparql lexer.
var Sparql = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "SPARQL",
		Aliases:   []string{"sparql"},
		Filenames: []string{"*.rq", "*.sparql"},
		MimeTypes: []string{"application/sparql-query"},
	},
	embedded,
	"embedded/sparql.xml",
))
