package h

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// HCL lexer.
var HCL = internal.Register(MustNewXMLLexer(
	&Config{
		Name:      "HCL",
		Aliases:   []string{"hcl"},
		Filenames: []string{"*.hcl"},
		MimeTypes: []string{"application/x-hcl"},
	},
	embedded,
	"embedded/hcl.xml",
))
