package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Solidity lexer.
var Solidity = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Solidity",
		Aliases:   []string{"sol", "solidity"},
		Filenames: []string{"*.sol"},
		MimeTypes: []string{},
		DotAll:    true,
	},
	embedded,
	"embedded/solidity.xml",
))
