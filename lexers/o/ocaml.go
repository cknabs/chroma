package o

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Ocaml lexer.
var Ocaml = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "OCaml",
		Aliases:   []string{"ocaml"},
		Filenames: []string{"*.ml", "*.mli", "*.mll", "*.mly"},
		MimeTypes: []string{"text/x-ocaml"},
	},
	embedded,
	"embedded/ocaml.xml",
))
