package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Elixir lexer.
var Elixir = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Elixir",
		Aliases:   []string{"elixir", "ex", "exs"},
		Filenames: []string{"*.ex", "*.exs"},
		MimeTypes: []string{"text/x-elixir"},
	},
	embedded,
	"embedded/elixir.xml",
))
