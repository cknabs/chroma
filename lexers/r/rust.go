package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Rust lexer.
var Rust = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "Rust",
		Aliases:   []string{"rust", "rs"},
		Filenames: []string{"*.rs", "*.rs.in"},
		MimeTypes: []string{"text/rust", "text/x-rust"},
		EnsureNL:  true,
	},
	embedded,
	"embedded/rust.xml",
))
