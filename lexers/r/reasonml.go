package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Reasonml lexer.
var Reasonml = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "ReasonML",
		Aliases:   []string{"reason", "reasonml"},
		Filenames: []string{"*.re", "*.rei"},
		MimeTypes: []string{"text/x-reasonml"},
	},
	embedded,
	"embedded/reasonml.xml",
))
