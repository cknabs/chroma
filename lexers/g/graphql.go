package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Go lexer.
var Graphql = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "GraphQL",
		Aliases:   []string{"graphql", "graphqls", "gql"},
		Filenames: []string{"*.graphql", "*.graphqls"},
	},
	embedded,
	"embedded/graphql.xml",
))
