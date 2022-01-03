package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Sieve lexer.
var Sieve = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Sieve",
		Aliases:   []string{"sieve"},
		Filenames: []string{"*.siv", "*.sieve"},
		MimeTypes: []string{},
	},
	embedded,
	"embedded/sieve.xml",
))
