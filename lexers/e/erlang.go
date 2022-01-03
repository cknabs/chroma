package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Erlang lexer.
var Erlang = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Erlang",
		Aliases:   []string{"erlang"},
		Filenames: []string{"*.erl", "*.hrl", "*.es", "*.escript"},
		MimeTypes: []string{"text/x-erlang"},
	},
	embedded,
	"embedded/erlang.xml",
))
